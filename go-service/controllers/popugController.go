package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "go-service/models"
    "log"
)

func CreatePopug(c *gin.Context, db *gorm.DB) {
    var json models.PopugModel

    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    go func() {
        result := db.Create(&json)
        if result.Error != nil {
            log.Println("Error creating record: ", result.Error)
        }
    }()
    
    c.Status(http.StatusAccepted)
}