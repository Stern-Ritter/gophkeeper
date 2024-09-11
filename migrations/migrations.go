package migrations

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed "*.sql"
var Migrations embed.FS

func MigrateDatabase(databaseDsn string) error {
	goose.SetBaseFS(Migrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose failed to set postgres dialect: %w", err)
	}

	db, err := goose.OpenDBWithDriver("pgx", databaseDsn)
	if err != nil {
		return fmt.Errorf("goose failed to open database connection: %w", err)
	}

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("goose failed to migrate database: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("goose failed to close database connection: %w", err)
	}

	return nil
}
