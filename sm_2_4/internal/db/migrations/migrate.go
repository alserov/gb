package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func MustMigrate(conn *sqlx.DB) {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		panic("failed to init driver: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres",
		driver)
	if err != nil {
		panic("failed to ini migrate: " + err.Error())
	}

	if err = m.Up(); err != nil && !errors.Is(migrate.ErrNoChange, err) {
		panic("failed to migrate: " + err.Error())
	}
}
