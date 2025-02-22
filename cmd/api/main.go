package main

import (
	"log"
	"os"

	"github.com/draysams/gosocial/internal/db"
	"github.com/draysams/gosocial/internal/env"
	"github.com/draysams/gosocial/internal/store"
)

func main() {
	cfg := config{
		port: env.GetString("PORT", ":8080"),
		db: dbConfig{
			connectionString:   env.GetString("DB_CONNECTION_STRING", "postgres://<admin>:<adminpswrd>@localhost:5432/<gosocial_db>?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.db.connectionString, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PORT")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
