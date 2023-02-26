package wishlists

import (
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
    routes.POST("/", h.CreateWishlist)
    routes.GET("/", h.GetWishlists)
    routes.GET("/:wishlist_id", h.GetWishlist)
    routes.DELETE("/:wishlist_id", h.ArchiveWishlist)

    routes.POST("/:wishlist_id/items", h.AddItem)
    routes.PUT("/:wishlist_id/items/:item_id", h.UpdateItem)
    routes.DELETE("/:wishlist_id/items/:item_id", h.RemoveItem)
}
