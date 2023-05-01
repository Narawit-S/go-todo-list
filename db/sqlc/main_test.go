package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDiver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/todo?sslmode=disable"
)

var testQuries *Queries

func TestMain(m *testing.M)()  {
	db, err := sql.Open(dbDiver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQuries = New(db)

	os.Exit((m.Run()))
}
