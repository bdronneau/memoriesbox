package db

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/bdronneau/memoriesbox/pkg/logger"

	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	DB *sql.DB

	logger *zap.SugaredLogger
}

type Config struct {
	dbHost     *string
	dbUser     *string
	dbName     *string
	dbPassword *string
	dbPort     *int
}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{
		dbHost:     fs.String("db-host", "0.0.0.0", "DB Hostname"),
		dbUser:     fs.String("db-user", "memoriesbox", "DB User"),
		dbName:     fs.String("db-name", "memoriesbox", "DB Name"),
		dbPassword: fs.String("db-password", "memoriesbox", "DB Password"),
		dbPort:     fs.Int("db-port", 5432, "DB Port"),
	}
}

func New(config Config, loggerApp logger.App) App {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", *config.dbUser, *config.dbPassword, *config.dbHost, *config.dbPort, *config.dbName)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		loggerApp.Sugar.Errorf("Unable to connect to database: %v", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		loggerApp.Sugar.Errorf("Unable to ping to database: %v", err)
		os.Exit(1)
	}

	loggerApp.Sugar.Info("Database connected")

	return App{
		DB: db,

		logger: loggerApp.Sugar,
	}
}
