package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initdb() {
	// Connect to the PostgreSQL database
	var err error

	DB, err = sql.Open("postgres", "postgres://postgres:2345@localhost/foomdies?sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database")
}
