package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"example/bucket/app/handlers"
	"example/bucket/app/middleware"
	"example/bucket/config"
	"example/bucket/config/db"
	"example/bucket/config/db/migration"
)

func init() {
	config.LoadEnvVariables()
	db.Init()
	migration.Migration(db.DB)
}

func setupRouter() *gin.Engine {
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
	return router
}

func main() {
	router := setupRouter()
	router.Run(":" + os.Getenv("PORT"))
}
