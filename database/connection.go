package database

import (
	"log"
	"medium_api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// connection to database
	connection, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	} else {
		log.Println("Database connected")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
