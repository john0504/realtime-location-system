package db

import (
	"database/sql"
	"log"

	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected")
}
