package main

import (
	"flag"
	"os"
	"scrapeit/app"
)

func main() {
	var (
		migrateDown bool
	)

	flag.BoolVar(&migrateDown, "migrate-down", false, "run database migrations down")
	flag.Parse()

	a := app.NewApp()
	// run db migration
	if migrateDown {
		a.MigrateDown()
		os.Exit(0)
	}
	a.MigrateUp()
	// start server
	a.Start()
}
