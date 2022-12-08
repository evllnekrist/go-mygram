package main

import (
	"log"
    "github.com/joho/godotenv"
	"go-rest-api/database"
	"go-rest-api/routers"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
