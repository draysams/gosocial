package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/draysams/gosocial/internal/env"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func run(direction string) {
	db, err := sql.Open("postgres", env.GetString("DB_CONNECTION_STRING", "postgres://<admin>:<adminpswrd>@localhost:5432/<gosocial_db>?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///Users/I343281/projies/project-2025/go/gosocial/cmd/migrate/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Migrating %s...", direction)

	if direction == "down" {
		log.Println("Rolling back migrations...")
		err = m.Down()
	} else {
		log.Println("Applying migrations...")
		err = m.Up()
	}

	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No change in migrations")
		} else {
			log.Fatal(err)
		}
	}
}

func main() {
	direction := flag.String("direction", "up", "Specify the migration direction: up or down")
	flag.Parse()

	log.Printf("Starting migration %s...", *direction)

	if *direction != "up" && *direction != "down" {
		log.Fatal("Invalid direction. Use 'up' or 'down'.")
	}

	run(*direction)
}
