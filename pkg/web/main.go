package web

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"memoriesbox/pkg/logger"
	"memoriesbox/pkg/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// App of package
type App interface {
	Start()
	Shutdown(context.Context) error
}

type app struct {
	done          chan struct{}
	echo          *echo.Echo
	listenAddress string
	logger        *zap.SugaredLogger
	repositories  repositories.App
}

type Config struct {
	address *string
	debug   *bool
	port    *int
}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{
		address: fs.String("api-address", "localhost", "API address"),
		debug:   fs.Bool("api-debug", false, "API Debug mode"),
		port:    fs.Int("api-port", 1080, "API Port"),
	}
}

func New(config Config, fs embed.FS, loggerApp logger.App, repoApp repositories.App) App {
	port := *config.port
	done := make(chan struct{})

	if port == 0 {
		loggerApp.Sugar.Fatal("Can not run on port 0")
	}

	loggerApp.Sugar.Infof("api listen on %s:%d", *config.address, *config.port)

	app := &app{
		repositories:  repoApp,
		done:          done,
		listenAddress: fmt.Sprintf("%s:%d", *config.address, port),
		logger:        loggerApp.Sugar,
	}

	app.echo = app.ConfigureEcho(*config.debug, fs)

	return app
}

func (a *app) Done() <-chan struct{} {
	return a.done
}

func (a *app) Start() {
	if err := a.echo.Start(a.listenAddress); err != nil && err != http.ErrServerClosed {
		a.echo.Logger.Fatalf("shutting down api server: ", err)
	}
}

// Shutdown gracefully shutdown HTTP
func (a *app) Shutdown(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}

func (a *app) ConfigureEcho(debug bool, embedFs fs.FS) *echo.Echo {
	e := echo.New()
	e.Debug = debug
	e.HidePort = true
	e.HideBanner = true

	tpl, err := template.ParseFS(embedFs, "templates/*.html")
	if err != nil {
		a.logger.Fatal("Can not read templates/*.html")
	}

	renderer := &TemplateRenderer{
		templates: tpl,
	}
	e.Renderer = renderer

	fsys, err := fs.Sub(embedFs, "static")
	if err != nil {
		a.logger.Fatal("Can not read ")
	}

	e.GET("/probes/ready", a.readyHandler)
	e.GET("/probes/live", a.liveHandler)
	e.GET("/version", a.versionHandler)

	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys)))))

	e.GET("/api/memories/count", a.countMemories)

	e.GET("/", a.getMemories)

	return e
}
