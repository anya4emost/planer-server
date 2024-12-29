package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func Connect(dbUrl string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		pqErr := err.(*pq.Error)
		log.Fatalf("Error connecting to database: %v\n", pqErr.Code)
	}

	log.Println("Connected to database")

	return db
}
