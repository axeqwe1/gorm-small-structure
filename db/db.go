package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to database")
	}
	DB = database
}
