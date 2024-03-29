package repositories

import (
	"context"
	"flag"
	"log/slog"
	"time"

	"github.com/bdronneau/memoriesbox/pkg/db"
	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/bdronneau/memoriesbox/pkg/repositories/models"
)

//go:generate mockgen -source main.go -destination ../mocks/repositories.go -mock_names App=Repositories -package mocks

// App of package
type App interface {
	AddMemory(quote string, author string, date time.Time) error
	CountMemories() (int64, error)
	GetRandomMemories() (models.Memory, error)
	PingDB() error
}

type app struct {
	dbApp db.App

	ExtraLog bool
}

type Config struct{}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{}
}

func New(config Config, loggerApp logger.App, dbApp db.App) App {
	return &app{
		dbApp:    dbApp,
		ExtraLog: loggerApp.ExtraLog,
	}
}

func (a *app) PingDB() error {
	err := a.dbApp.Ping(context.Background())
	if err != nil {
		slog.Error("Unable to ping to database", "error", err)
		return err
	}

	return nil
}
