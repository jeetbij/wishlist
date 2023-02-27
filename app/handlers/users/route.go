package users

import (
	"example/bucket/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/users")
	routes.POST("/signup", h.SignUp)
	routes.POST("/login", h.LogIn)
	routes.GET("/validate", middleware.RequireAuth, h.Validate)
}
