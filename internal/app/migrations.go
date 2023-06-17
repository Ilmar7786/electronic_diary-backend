package app

import (
	"flag"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

const (
	defaultValue = "disable"
	flagName     = "migrations"
)

func (a *App) migrations() {
	var command string
	flag.StringVar(&command, flagName, defaultValue, "this is app config file")
	flag.Parse()

	if strings.ToLower(command) == defaultValue {
		return
	}

	db, err := a.pgClient.DB()
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)

	switch strings.ToLower(command) {
	case "up":
		log.Println("run up migrations")
		if err := m.Up(); err != nil {
			log.Fatalln(err)
		}
	case "down":
		log.Println("run down migrations")
		if err := m.Down(); err != nil {
			log.Fatalln(err)
		}
	case "":
	default:
		log.Println("wrong command")
	}
}
