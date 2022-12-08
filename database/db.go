package database

import (
	// "os"
	"fmt"
	"log"
	"go-rest-api/models"
	"gorm.io/driver/postgres"
    "github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	err      error
)

func StartDB() {
    var envs map[string]string
    envs, err := godotenv.Read(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	host     	:= envs["DB_HOST"]
	user 		:= envs["DB_USER"]
	password 	:= envs["DB_PASSWORD"]
	dbPort   	:= envs["DB_PORT"]
	dbName   	:= envs["DB_DATABASE"]

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	fmt.Println("successfully connected to database")
	db.AutoMigrate(&models.User{},&models.Photo{},&models.Comment{},&models.Socialmedia{})
}

func GetDB() *gorm.DB {
	return db
}