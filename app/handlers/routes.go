package handlers

import (
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
	wishlists.RegisterRoutes(router, db)
	users.RegisterRoutes(router, db)
}
