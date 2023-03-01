package wishlists

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"example/bucket/app/helpers"
	"example/bucket/app/models/wishlist"
)

func CreateWishlist(ctx *gin.Context) {
	token := helpers.GetGuestToken(ctx)
	usr, _ := helpers.GetUser(ctx)
	userId := helpers.GetUserId(usr)
	body := helpers.CreateWishlistRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var wishlst wishlist.Wishlist

	wishlst.UserId = userId
	wishlst.Name = body.Name
	wishlst.Type = body.Type
	wishlst.Description = body.Description
	wishlst.Token = token

	if result := wishlist.DB().Create(&wishlst); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &wishlst)
}

func GetWishlists(ctx *gin.Context) {
	token := helpers.GetGuestToken(ctx)
	usr, _ := helpers.GetUser(ctx)

	var wishlists []wishlist.Wishlist
	var result *gorm.DB
	if usr.ID == 0 {
		result = wishlist.GuestWishlists(token).Find(&wishlists)
	} else {
		result = wishlist.UserWishlists(usr.ID).Find(&wishlists)
	}

	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &wishlists)
}

func GetWishlist(ctx *gin.Context) {
	token := helpers.GetGuestToken(ctx)
	usr, _ := helpers.GetUser(ctx)
	wishlistId := ctx.Param("wishlist_id")

	var wishlst wishlist.Wishlist
	var result *gorm.DB
	if usr.ID == 0 {
		result = wishlist.GuestWishlists(token).Where("id = ?", wishlistId).First(&wishlst)
	} else {
		result = wishlist.UserWishlists(usr.ID).Where("id = ?", wishlistId).First(&wishlst)
	}

	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &wishlst)
}

func ArchiveWishlist(ctx *gin.Context) {
	token := helpers.GetGuestToken(ctx)
	usr, _ := helpers.GetUser(ctx)
	wishlistId := ctx.Param("wishlist_id")

	var wishlst wishlist.Wishlist
	var result *gorm.DB
	if usr.ID == 0 {
		result = wishlist.GuestWishlists(token).First(&wishlst, wishlistId)
	} else {
		result = wishlist.UserWishlists(usr.ID).First(&wishlst, wishlistId)
	}

	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	wishlst.Archived = true
	wishlist.DB().Save(&wishlst)

	ctx.Status(http.StatusOK)
}
