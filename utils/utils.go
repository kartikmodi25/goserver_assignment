package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		log.Fatal(err)
	}

	// Create user table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			dob DATE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create movie table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS movies (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			title VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("DB STARTED")
	return db
}
