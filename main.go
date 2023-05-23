package main

import (
	"database/sql"
	"goserver2/router"
	"goserver2/services"
	"goserver2/utils"
	"log"

	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db = utils.GetConnection()
	services.SetDB(db)
	defer db.Close()
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
