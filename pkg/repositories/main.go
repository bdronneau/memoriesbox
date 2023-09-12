package repositories

import (
	"database/sql"
	"flag"
	"time"

	"github.com/bdronneau/memoriesbox/pkg/db"
	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/bdronneau/memoriesbox/pkg/repositories/models"

	"go.uber.org/zap"
)

//go:generate mockgen -source main.go -destination ../mocks/repositories.go -mock_names App=Repositories -package mocks

// App of package
type App interface {
	AddMemory(quote string, author string, date time.Time) error
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
