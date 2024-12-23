package routers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "go-service/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    router.POST("/popug", func(c *gin.Context) {
        controllers.CreatePopug(c, db)
    })

    return router
}