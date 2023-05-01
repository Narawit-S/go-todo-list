package main

import (
	"database/sql"
	"log"

	"github.com/Narawit-S/go-todo-list/api"
	db "github.com/Narawit-S/go-todo-list/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDiver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/todo?sslmode=disable"
	port = ":8080"
)

func main() {
	dbConn, err := sql.Open(dbDiver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(dbConn)
	server := api.NewServer(store)

	err = server.Start(port)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
