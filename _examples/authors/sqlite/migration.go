// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-connect).

package main

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed sql/migrations
var migrations embed.FS

func ensureSchema(db *sql.DB) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	return goose.Up(db, "sql/migrations")

}
