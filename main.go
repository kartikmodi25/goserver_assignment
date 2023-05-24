package main

import (
	"database/sql"
	"goserver2/router"
	"goserver2/services"
	"goserver2/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	db = utils.GetConnection()
	services.SetDB(db)
	defer db.Close()
	r := router.SetupRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid port number")
	}
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
}
