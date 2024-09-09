package web

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/bdronneau/memoriesbox/pkg/repositories"
	"gitlab.com/greyxor/slogor"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/labstack/echo/v4"
)

// App of package
type App interface {
	GetEcho() *echo.Echo
	Start()
	Shutdown(context.Context) error
}

type app struct {
	repositories  repositories.App
	done          chan struct{}
	echo          *echo.Echo
	listenAddress string

	ExtraLog bool

	tracerProvider *sdktrace.TracerProvider
}

type Config struct {
	address       *string
	debug         *bool
	port          *int
	featAddMemory *bool

	tracingName     *string
	tracingEnabled  *bool
	tracingOtelRate *string
}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{
		address:       fs.String("api-address", "localhost", "API address"),
		debug:         fs.Bool("api-debug", false, "API Debug mode"),
		port:          fs.Int("api-port", 1080, "API Port"),
		featAddMemory: fs.Bool("feat-add-memory", true, "Interface to add memory"),

		tracingEnabled:  fs.Bool("tracing-enabled", false, "Enable tracing"),
		tracingName:     fs.String("tracing-name", "memoriesbox", "Server name for tracing"),
		tracingOtelRate: fs.String("tracing-rate", "always", "Which rate for tracing"),
	}
}

func New(config Config, fs fs.FS, loggerApp logger.App, repoApp repositories.App) App {
	port := *config.port
	done := make(chan struct{})

	if port == 0 {
		slog.Error("Can not run on port 0")
		os.Exit(1)
	}

	slog.Info("api listen on", "address", fmt.Sprintf("%s:%d", *config.address, *config.port))

	app := &app{
		repositories:  repoApp,
		done:          done,
		listenAddress: fmt.Sprintf("%s:%d", *config.address, port),
		ExtraLog:      loggerApp.ExtraLog,
	}

	app.echo = app.ConfigureEcho(*config.debug, fs, *config.featAddMemory)

	if *config.tracingEnabled {
		tracerProvider, err := app.InitTracer(*config.tracingName, *config.tracingOtelRate)
		if err != nil {
			slog.Error("Can not init tracer", slogor.Err(err))
			os.Exit(1)
		}
		app.tracerProvider = tracerProvider
	}

	return app
}

func (a *app) Done() <-chan struct{} {
	return a.done
}

func (a *app) Start() {
	defer func() {
		if err := a.tracerProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	if err := a.echo.Start(a.listenAddress); err != nil && err != http.ErrServerClosed {
		a.echo.Logger.Fatalf("shutting down api server: ", err)
	}
}

// Shutdown gracefully shutdown HTTP
func (a *app) Shutdown(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}

func (a *app) ConfigureEcho(debug bool, embedFs fs.FS, featAddMemory bool) *echo.Echo {
	e := echo.New()
	e.Debug = debug
	e.HidePort = true
	e.HideBanner = true

	tpl, err := template.ParseFS(embedFs, "templates/partials/*.html", "templates/*.html")
	if err != nil {
		slog.Error("Can not read templates/*.html")
		os.Exit(1)
	}

	renderer := &TemplateRenderer{
		templates: tpl,
	}
	e.Renderer = renderer

	fsys, err := fs.Sub(embedFs, "static")
	if err != nil {
		slog.Error("Can not read", slogor.Err(err))
		os.Exit(1)
	}

	e.Use(otelecho.Middleware("memoriesbox"))

	e.GET("/probes/ready", a.readyHandler)
	e.GET("/probes/live", a.liveHandler)
	e.GET("/version", a.versionHandler)

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys)))))

	e.GET("/api/memories/count", a.countMemories)

	e.GET("/", a.getMemories)

	if featAddMemory {
		e.POST("/api/memories/add", a.addAPIMemory)
		e.GET("/add", a.addMemory)
	}

	return e
}

func (a *app) GetEcho() *echo.Echo {
	return a.echo
}

func (a *app) InitTracer(name string, sampler string) (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	otelServiceName, bool := os.LookupEnv("OTEL_SERVICE_NAME")
	if bool {
		name = otelServiceName
	}

	otelResource, err := newResource(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("resource: %w", err)
	}

	tracerExporter, err := newTraceExporter(ctx, os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		return nil, fmt.Errorf("trace exporter: %w", err)
	}

	otelServiceSampler, bool := os.LookupEnv("OTEL_TRACES_SAMPLER")
	if bool {
		sampler = otelServiceSampler
	}

	otelSampler, err := newSampler(strings.TrimSpace(sampler))
	if err != nil {
		return nil, fmt.Errorf("sampler: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(tracerExporter),
		trace.WithResource(otelResource),
		trace.WithSampler(otelSampler),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tracerProvider, nil
}

func newResource(ctx context.Context, serviceName string) (*resource.Resource, error) {
	newResource, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithAttributes(
			// semconv.ServiceVersion(model.Version()),
			// attribute.String("git.commit.sha", model.GitSha()),
			attribute.String("service.name", serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create resource: %w", err)
	}

	r, err := resource.Merge(resource.Default(), newResource)
	if err != nil {
		return nil, fmt.Errorf("merge resource with default: %w", err)
	}

	return r, nil
}

func newSampler(rate string) (trace.Sampler, error) {
	switch rate {
	case "always":
		return trace.AlwaysSample(), nil

	case "never":
		return trace.NeverSample(), nil

	default:
		rateRatio, err := strconv.ParseFloat(rate, 64)
		if err != nil {
			return nil, fmt.Errorf("parse sample rate `%s`: %w", rate, err)
		}

		return trace.TraceIDRatioBased(rateRatio), nil
	}
}

func newTraceExporter(ctx context.Context, endpoint string) (trace.SpanExporter, error) {
	return otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)
}
