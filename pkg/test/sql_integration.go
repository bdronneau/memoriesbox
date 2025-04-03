package test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	migratePgx "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresIntegration struct {
	t      *testing.T
	db     *sql.DB
	config DatabaseConfig
}

func NewIntegration(t *testing.T) *PostgresIntegration {
	t.Helper()

	return &PostgresIntegration{t: t}
}

func (pi *PostgresIntegration) Bootstrap(name string) {
	pi.config = getDatabaseConfig(pi.t)

	ctx := context.Background()

	pi.config.PostgresDatabase = "postgres"

	pi.connect(ctx)
	pi.recreateDatabase(ctx, name)
	pi.Close()

	pi.config.PostgresDatabase = name

	pi.connect(ctx)
	pi.migrate()
	pi.truncateAll(ctx)
}

func (pi *PostgresIntegration) connect(ctx context.Context) {
	var err error

	pi.db, err = sql.Open("pgx", pi.config.URI())
	if err != nil {
		pi.t.Fatal(err)
	}
}

func (pi *PostgresIntegration) recreateDatabase(ctx context.Context, name string) {
	if _, err := pi.db.ExecContext(ctx, fmt.Sprintf(`DROP DATABASE IF EXISTS "%s"`, name)); err != nil {
		pi.t.Fatal(err)
	}

	if _, err := pi.db.ExecContext(ctx, fmt.Sprintf(`CREATE DATABASE "%s"`, name)); err != nil {
		pi.t.Fatal(err)
	}
}

func (pi *PostgresIntegration) migrate() {
	db, err := sql.Open("pgx", pi.config.URI())
	if err != nil {
		pi.t.Fatal(fmt.Errorf("open pgx: %w", err))
	}
	pi.db = db

	migrateInstance, err := migratePgx.WithInstance(db, &migratePgx.Config{MultiStatementEnabled: true})
	if err != nil {
		pi.t.Fatal(fmt.Errorf("init migrate: %w", err))
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://../../db/migrations", "pgx", migrateInstance)
	if err != nil {
		pi.t.Fatal(fmt.Errorf("create migrator: %w", err))
	}

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		pi.t.Fatal(fmt.Errorf("migrate: %w", err))
	}
}

func (pi *PostgresIntegration) truncateAll(ctx context.Context) {
	if _, err := pi.db.ExecContext(ctx, truncateAllUserTablesFunction); err != nil {
		pi.t.Fatal(err)
	}
}

func (pi *PostgresIntegration) DB() *sql.DB {
	return pi.db
}

func (pi *PostgresIntegration) Close() {
	defer func() {
		if err := pi.db.Close(); err != nil {
			pi.t.Errorf("Error closing the database: %v", err)
		}
	}()

}

func (pi *PostgresIntegration) Reset(ctx context.Context) {
	query := fmt.Sprintf(`SELECT truncate_tables('%s');`, pi.config.PostgresUsername)
	if _, err := pi.db.ExecContext(ctx, query); err != nil {
		pi.t.Fatal(err)
	}
}

const truncateAllUserTablesFunction = `
CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
DECLARE
    statements CURSOR FOR
        SELECT schemaname, tablename FROM pg_tables
        WHERE tableowner = username
          AND schemaname NOT IN ('pg_catalog', 'information_schema')
          AND tablename != 'market';
BEGIN
    FOR stmt IN statements LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.schemaname) || '.' || quote_ident(stmt.tablename) || ' CASCADE;';
    END LOOP;
END;
$$ LANGUAGE plpgsql;`

type DatabaseConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDatabase string
	PostgresUsername string
	PostgresPassword string
}

func getDatabaseConfig(t *testing.T) DatabaseConfig {
	return DatabaseConfig{
		PostgresHost:     getEnvWithDefault("MEMORIESBOX_TEST_DB_HOST", "127.0.0.1"),
		PostgresPort:     getEnvWithDefault("MEMORIESBOX_TEST_DB_PORT", "5432"),
		PostgresDatabase: getEnvWithDefault("MEMORIESBOX_TEST_DB_NAME", "memoriesbox"),
		PostgresUsername: getEnvWithDefault("MEMORIESBOX_TEST_DB_USER", "postgres"),
		PostgresPassword: getEnvWithDefault("MEMORIESBOX_TEST_DB_PASS", "pwd"),
	}
}

func (c DatabaseConfig) URI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&application_name=%s",
		c.PostgresUsername, c.PostgresPassword, c.PostgresHost, c.PostgresPort, c.PostgresDatabase, "memoriesbox_test")
}

func getEnvWithDefault(name, defaultValue string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	return defaultValue
}
