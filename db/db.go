package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connectionStr := "user=postgres dbname=alura_loja password=1540 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
