package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"example/bucket/app/config"
	"example/bucket/app/config/db"
	"example/bucket/app/config/db/migration"
	"example/bucket/app/handlers"
	"example/bucket/app/middleware"
)

func init() {
	config.LoadEnvVariables()
	db.Init()
	migration.Migration(db.DB)
}

func main() {
	if os.Getenv("MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	// Add Middlewares
	router.Use(middleware.SetGuestToken)
	router.Use(middleware.RequireAuth)

	// Register Routes
	handlers.RegisterRoutes(router, db.DB)

	router.Run(":" + os.Getenv("PORT"))
}
