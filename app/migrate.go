package app

import (
	"embed"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (a App) MigrateDown() {
	defer a.db.Close()
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(a.db, "migrations"); err != nil {
		panic(err)
	}
}

func (a App) MigrateUp() {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(a.db, "migrations"); err != nil {
		panic(err)
	}
}
