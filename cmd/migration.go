package main

import (
	"database/sql"
	"fmt"
	"tethys-go/internal/core/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func migration() error {
	fmt.Println(config.Get().PgConfig.GetDSN())

	db, err := sql.Open("postgres", config.Get().PgConfig.GetDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	return goose.Up(db, "/app/migrations")
}
