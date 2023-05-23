package main

import (
	"database/sql"
	"goserver2/router"
	"goserver2/services"
	"goserver2/utils"
	"log"
)

var db *sql.DB

func main() {
	db = utils.GetConnection()
	services.SetDB(db)
	defer db.Close()
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
