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
	// fmt.Println(db)
	services.SetDB(db)
	defer db.Close()
	// fmt.Println(db)
	// var err error
	// db, err = sql.Open("postgres", "postgres://postgres:252900@localhost/postgres2?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//

	// Setup router;
	r := router.SetupRouter()
	// fmt.Println(r)
	// log.Println("Listening on Port 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
	// // // Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
