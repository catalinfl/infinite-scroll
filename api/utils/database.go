package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Database *sql.DB

func ConnectDB() {
	godotenv.Load()

	url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	Database = db

}
