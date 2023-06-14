package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func OpenPostgresDB() *sql.DB {
	connect := "user=mishar password=1234 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
