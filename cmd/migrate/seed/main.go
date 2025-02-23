package main

import (
	"github.com/draysams/gosocial/internal/db"
	"github.com/draysams/gosocial/internal/env"
	"github.com/draysams/gosocial/internal/store"
)

func main() {
	connectionString := env.GetString("DB_CONNECTION_STRING", "postgres://postgres:postgres@localhost:5432/gosocial?sslmode=disable")
	connection, err := db.New(connectionString, 30, 30, "30m")
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	store := store.NewStorage(connection)

	db.Seed(store)

}
