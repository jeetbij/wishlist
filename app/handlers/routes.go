package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example/bucket/app/handlers/users"
	"example/bucket/app/handlers/wishlists"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	router.GET("/health_check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
	router.GET("/register", func(ctx *gin.Context) {
		log.Println("[EXTENTION_INSTALLED]")
		ctx.JSON(200, gin.H{
			"status": "OK",
		})
	})
	users.RegisterRoutes(router, db)
	wishlists.RegisterRoutes(router, db)
}
