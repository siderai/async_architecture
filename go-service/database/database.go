package database

import (
	"go-service/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	if err := db.AutoMigrate(&models.PopugModel{}); err != nil {
		log.Fatal("failed to migrate database")
	}

	return db
}
