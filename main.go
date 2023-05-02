package main

import (
	"database/sql"
	"log"

	"github.com/Narawit-S/go-todo-list/api"
	db "github.com/Narawit-S/go-todo-list/db/sqlc"
	"github.com/Narawit-S/go-todo-list/utils"
	_ "github.com/lib/pq"
)

func main() {
	env, err := utils.LoadEnv(".")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}

	dbConn, err := sql.Open(env.DBDriver, env.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(dbConn)
	server := api.NewServer(store)

	err = server.Start(env.Port)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
