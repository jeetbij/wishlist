package users

import (
	"example/bucket/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	routes := router.Group("/users")
	routes.POST("/signup", SignUp)
	routes.POST("/login", LogIn)
	routes.GET("/validate", middleware.RequireAuth, Validate)
}
