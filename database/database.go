package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
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

	// whether we are running a test database or a production (which is in the cloud)
	// will use different files to start the db
	environment, _ := os.LookupEnv("ENVIRONMENT")
	var migrationPath string
	// Runs when we use docker (this runs in the cloud either on the dev side or productin side)
	migrationPath = "file://database/migrations"
	if environment == "production" || environment == "dev" {
		fmt.Println("Using dev migrations")
		migrationPath = "file://database/migrations"
	} else {
		// This runs when we run our tests (this is run in the docker test files)
		fmt.Println("Using test migrations")
		migrationPath = "file://database/test-migrations"
	}
	m, err := migrate.NewWithDatabaseInstance(
		// for production
		migrationPath,

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
