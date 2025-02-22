package main

import (
	"log"
	"os"

	"github.com/draysams/gosocial/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	os.LookupEnv("PORT")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
