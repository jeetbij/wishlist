package wishlists

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"example/bucket/app/helpers"
	"example/bucket/app/models/wishlist"
)

func (h handler) CreateWishlist(ctx *gin.Context) {
	usr, _ := helpers.GetUser(ctx)
	userId := helpers.GetUserId(usr)
	body := helpers.CreateWishlistRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var wishlist wishlist.Wishlist

	wishlist.UserId = userId
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
	var wishlists []wishlist.Wishlist

	if result := h.DB.Scopes(wishlist.UnarchivedWishlist).Preload("Items", "is_active = ?", true).Find(&wishlists); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &wishlists)
}

func (h handler) GetWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlist_id")

	var wishlst wishlist.Wishlist

	result := h.DB.Scopes(wishlist.UnarchivedWishlist).Preload("Items", "is_active = ?", true).Where("id = ?", wishlistId).First(&wishlst)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &wishlst)
}

func (h handler) ArchiveWishlist(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlist_id")

	var wishlst wishlist.Wishlist

	result := h.DB.Scopes(wishlist.UnarchivedWishlist).First(&wishlst, wishlistId)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	wishlst.Archived = true
	h.DB.Save(&wishlst)

	ctx.Status(http.StatusOK)
}
