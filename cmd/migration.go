package main

import (
	"tethys-go/internal/core/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func migration() error {
	db, err := goose.OpenDBWithDriver("postgres", config.Get().PgConfig.GetDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	return goose.Up(db, "migrations")
}
