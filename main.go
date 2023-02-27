package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"example/bucket/app/common/config"
	"example/bucket/app/common/db"
	"example/bucket/app/handlers"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	dbHandler := db.Init()

	db.Migration(dbHandler)

	handlers.RegisterRoutes(router, dbHandler)

	router.Run(os.Getenv("PORT"))
}
