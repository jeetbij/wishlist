package main

import (
	"github.com/gin-gonic/gin"

	"example/bucket/app/common/db"
	"example/bucket/app/handlers"
)

func main() {
	var port string = ":3000"
	router := gin.Default()
	dbHandler := db.Init()

	handlers.RegisterRoutes(router, dbHandler)

	router.Run(port)
}
