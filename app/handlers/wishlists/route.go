package wishlists

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

	routes := router.Group("/wishlists")
	routes.POST("/", middleware.RequireAuth, h.CreateWishlist)
	routes.GET("/", middleware.RequireAuth, h.GetWishlists)
	routes.GET("/:wishlist_id", middleware.RequireAuth, h.GetWishlist)
	routes.DELETE("/:wishlist_id", middleware.RequireAuth, h.ArchiveWishlist)

	routes.POST("/:wishlist_id/items", middleware.RequireAuth, h.AddItem)
	routes.PUT("/:wishlist_id/items/:item_id", middleware.RequireAuth, h.UpdateItem)
	routes.DELETE("/:wishlist_id/items/:item_id", middleware.RequireAuth, h.RemoveItem)
}
