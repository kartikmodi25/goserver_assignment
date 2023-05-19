package main

import (
	"database/sql"
	"goserver2/router"
	"log"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup router
	r := router.SetupRouter()

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
