package main

import (
	"go-service/database"
	"go-service/routers"
	"log"
)

func main() {
	log.Println("Starting")
	db := database.SetupDatabase()
	routers.SetupRouter(db).Run(":8080")
	log.Println("Started")
}
