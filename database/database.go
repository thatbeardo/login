package database

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func EstablishConnection(dbURL string) (*sql.DB, error) {
	c, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("parsing postgres URI: %w", err)
	}

	db := stdlib.OpenDB(*c)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = validateSchema(db)
	if err != nil {
		return nil, fmt.Errorf("validate schema failed: %w", err)
	}

	return db, nil
}

// Migrate migrates the Postgres schema to the current version.
func validateSchema(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../../database/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
