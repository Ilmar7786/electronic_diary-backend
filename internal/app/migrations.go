package app

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	commandUp   = "up"
	commandDown = "down"
)

func (a *App) migrations(option string) {
	sql, err := a.pgClient.DB()
	if err != nil {
		log.Fatalln(err)
	}

	instance, err := postgres.WithInstance(sql, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		instance,
	)
	if err != nil {
		log.Fatal(err)
	}

	switch option {
	case commandUp:
		log.Println("migration starting up")
		err := m.Up()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("migration success up")
	case commandDown:
		log.Println("migration starting down")
		err := m.Down()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("migration success down")
	case "":
	default:
		log.Fatalln("invalid command")

	}
}
