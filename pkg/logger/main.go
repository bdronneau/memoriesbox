package logger

import (
	"flag"
	"log/slog"
	"os"
)

type App struct {
	Logger *slog.Logger

	ExtraLog bool
}

type Config struct {
	logDev        *bool
	logExtraDebug *bool
	LogLevel      *string
}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{
		logDev:        fs.Bool("log-dev", false, "display in dev mode"),
		logExtraDebug: fs.Bool("log-extra-debug", false, "Always more"),
		LogLevel:      fs.String("log-level", "info", "Change log level"),
	}
}

func New(config Config) App {
	var level slog.Level

	if err := level.UnmarshalText([]byte(*config.LogLevel)); err != nil {
		slog.Error(err.Error(), "level", *config.LogLevel)

		return App{
			ExtraLog: *config.logExtraDebug,
		}
	}

	options := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler
	if *config.logDev {
		handler = slog.NewTextHandler(os.Stdout, options)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, options)
	}

	slog.SetDefault(slog.New(handler))

	return App{
		ExtraLog: *config.logExtraDebug,
	}
}
