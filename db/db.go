package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func OpenPostgresDB() (*sql.DB, error) {
	connect := "user=mishar password=1234 dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	return db, nil
}
