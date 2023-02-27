package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example/bucket/app/handlers/users"
	"example/bucket/app/handlers/wishlists"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	wishlists.RegisterRoutes(router, db)
	users.RegisterRoutes(router, db)
}
