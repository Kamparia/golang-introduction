package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DbConnect() *sqlx.DB {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to postgres database
	db, err := sqlx.Connect("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
