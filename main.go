package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"example/bucket/app/config"
	"example/bucket/app/config/db"
	"example/bucket/app/config/db/migration"
	"example/bucket/app/handlers"
)

func init() {
	config.LoadEnvVariables()
	db.Init()
	migration.Migration(db.DB)
}

func main() {
	router := gin.Default()

	handlers.RegisterRoutes(router, db.DB)

	router.Run(os.Getenv("PORT"))
}
