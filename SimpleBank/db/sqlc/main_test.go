package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

var err error

func TestMain(m *testing.M) {
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("comm.sccount to db:", err)
	}

	testQueries = New(testDB)

	m.Run()

	os.Exit(m.Run())
}
