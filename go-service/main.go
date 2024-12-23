package main

import (
    "go-service/database"
    "go-service/routers"
)

func main() {
    db := database.SetupDatabase()
    routers.SetupRouter(db).Run(":8080")
}