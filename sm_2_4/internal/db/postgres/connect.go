package postgres

import (
	"github.com/alserov/gb/sm_2_4/internal/db/migrations"
	"github.com/jmoiron/sqlx"
)

func MustConnect(dsn string) *sqlx.DB {
	conn := sqlx.MustConnect("postgres", dsn)

	migrations.MustMigrate(conn)

	return conn
}
