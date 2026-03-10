package dbconfig

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB(databaseUrl string) *sql.DB {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatalf("Error connecting to the database. Reason: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed. Reason: %v", err)
	}

	log.Println("Connected to Database")

	return db
}
