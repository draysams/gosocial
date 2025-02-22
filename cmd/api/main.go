package main

import (
	"log"
	"os"

	"github.com/draysams/gosocial/internal/env"
	"github.com/draysams/gosocial/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	os.LookupEnv("PORT")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
