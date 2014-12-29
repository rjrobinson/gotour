package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgress@localhost/godos_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var subject string

	rows, err := db.Query("select subject from todos")
	for rows.Next() {
		rows.Scan(&subject)
	}
}
