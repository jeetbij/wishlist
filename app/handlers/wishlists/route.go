package wishlists

import (
	"example/bucket/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	routes := router.Group("/wishlists")
	routes.POST("/", middleware.RequireAuth, CreateWishlist)
	routes.GET("/", middleware.RequireAuth, GetWishlists)
	routes.GET("/:wishlist_id", middleware.RequireAuth, GetWishlist)
	routes.DELETE("/:wishlist_id", middleware.RequireAuth, ArchiveWishlist)

	routes.POST("/:wishlist_id/items", middleware.RequireAuth, AddItem)
	routes.PUT("/:wishlist_id/items/:item_id", middleware.RequireAuth, UpdateItem)
	routes.DELETE("/:wishlist_id/items/:item_id", middleware.RequireAuth, RemoveItem)
}
