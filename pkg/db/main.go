package db

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/bdronneau/memoriesbox/pkg/logger"

	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	ErrNoHost  = errors.New("no host for database connection")
	SQLTimeout = time.Second * 5
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
	dbSsl      *string
}

func GetConfig(fs *flag.FlagSet) Config {
	return Config{
		dbHost:     fs.String("db-host", "0.0.0.0", "DB Hostname"),
		dbUser:     fs.String("db-user", "memoriesbox", "DB User"),
		dbName:     fs.String("db-name", "memoriesbox", "DB Name"),
		dbSsl:      fs.String("db-ssl", "prefer", "DB SSLMode"),
		dbPassword: fs.String("db-password", "memoriesbox", "DB Password"),
		dbPort:     fs.Int("db-port", 5432, "DB Port"),
	}
}

func New(config Config, loggerApp logger.App) (App, error) {
	host := strings.TrimSpace(*config.dbHost)
	if len(host) == 0 {
		return App{}, ErrNoHost
	}

	user := strings.TrimSpace(*config.dbUser)
	pass := *config.dbPassword
	name := strings.TrimSpace(*config.dbName)
	sslmode := *config.dbSsl

	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&application_name=%s", user, pass, host, *config.dbPort, name, sslmode, "memoriesbox"))
	if err != nil {
		return App{}, fmt.Errorf("connect to postgres: %w", err)
	}

	instance := App{
		DB: db,

		logger: loggerApp.Sugar,
	}

	ctx, cancel := context.WithTimeout(context.Background(), SQLTimeout)
	defer cancel()

	return instance, instance.Ping(ctx)
}

func (a App) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, SQLTimeout)
	defer cancel()

	return a.DB.PingContext(ctx)
}
