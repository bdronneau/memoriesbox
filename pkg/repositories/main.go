package repositories

import (
	"database/sql"
	"flag"
	"memoriesbox/pkg/db"
	"memoriesbox/pkg/logger"
	"memoriesbox/pkg/repositories/models"

	"go.uber.org/zap"
)

// App of package
type App interface {
	CountMemories() int64
	GetRandomMemories() (models.Memory, error)
	PingDB() error
}

type app struct {
	DB *sql.DB

	logger *zap.SugaredLogger
}

type Config struct{}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{}
}

func New(config Config, loggerApp logger.App, dbApp db.App) App {
	return &app{
		DB:     dbApp.DB,
		logger: loggerApp.Sugar,
	}
}

func (a *app) PingDB() error {
	err := a.DB.Ping()
	if err != nil {
		a.logger.Errorf("Unable to ping to database: %v", err)
		return err
	}

	return nil
}
