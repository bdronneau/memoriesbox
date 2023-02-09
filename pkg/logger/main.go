package logger

import (
	"flag"
	"log"

	"go.uber.org/zap"
)

type App struct {
	Sugar *zap.SugaredLogger
}

func New(fs *flag.FlagSet) App {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	// flushes buffer, if any
	defer func() { _ = logger.Sync() }()

	return App{
		Sugar: logger.Sugar(),
	}
}
