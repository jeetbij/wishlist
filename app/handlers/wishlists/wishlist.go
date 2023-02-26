package wishlists

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example/bucket/app/models"
	"example/bucket/app/handlers/wishlists/helpers"
)

func (h handler) CreateWishlist(ctx *gin.Context) {
	body := helpers.CreateWishlistRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
        return
	}

	var wishlist models.Wishlist

	wishlist.Name = body.Name
	wishlist.Type = body.Type
	wishlist.Description = body.Description

	if result := h.DB.Create(&wishlist); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
	}

	ctx.JSON(http.StatusCreated, &wishlist)
}

func (h handler) GetWishlists(ctx *gin.Context) {
    var wishlists []models.Wishlist

    if result := h.DB.Scopes(models.UnarchivedWishlist).Preload("Items", "is_active = ?", true).Find(&wishlists); result.Error != nil {
		log.Println(result.Error)
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &wishlists)
}

func (h handler) GetWishlist(ctx *gin.Context) {
    wishlistId := ctx.Param("wishlist_id")

    var wishlist models.Wishlist

    result := h.DB.Scopes(models.UnarchivedWishlist).Preload("Items", "is_active = ?", true).Where("id = ?", wishlistId).First(&wishlist)
	if result.Error != nil {
		log.Println(result.Error)
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    ctx.JSON(http.StatusOK, &wishlist)
}

func (h handler) ArchiveWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlist_id")

	var wishlist models.Wishlist

    result := h.DB.Scopes(models.UnarchivedWishlist).First(&wishlist, wishlistId)
	if result.Error != nil {
		log.Println(result.Error)
        ctx.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

	wishlist.Archived = true
	h.DB.Save(&wishlist)

    ctx.Status(http.StatusOK)
}
