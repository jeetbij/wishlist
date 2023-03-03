package wishlists

import (
	"example/bucket/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	routes := router.Group("/wishlists")
	routes.Use(middleware.SetGuestToken)

	routes.POST("/", CreateWishlist)
	routes.GET("/", GetWishlists)
	routes.GET("/:wishlist_id", GetWishlist)
	routes.DELETE("/:wishlist_id", ArchiveWishlist)

	routes.POST("/:wishlist_id/items", AddItem)
	routes.PUT("/:wishlist_id/items/:item_id", UpdateItem)
	routes.DELETE("/:wishlist_id/items/:item_id", RemoveItem)
}
