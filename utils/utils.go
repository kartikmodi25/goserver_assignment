package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:252900@localhost/postgres2?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

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
