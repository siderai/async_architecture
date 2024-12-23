package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "go-service/models"
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