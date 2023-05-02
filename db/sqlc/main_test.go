package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Narawit-S/go-todo-list/utils"
	_ "github.com/lib/pq"
)

var testQuries *Queries

func TestMain(m *testing.M)()  {
	env, err := utils.LoadEnv("../..")
	if err != nil {
		log.Fatal("Cannot load env", err)
	}

	db, err := sql.Open(env.DBDriver, env.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQuries = New(db)

	os.Exit((m.Run()))
}
