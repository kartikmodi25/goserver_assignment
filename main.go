package main

import (
	"fmt"
	"goserver/utils"
	"log"
)

func main() {
	log.Println("Main app started")
	var dbConn = utils.GetConnection()
	fmt.Println(dbConn)
}
