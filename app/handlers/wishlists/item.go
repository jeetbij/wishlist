package wishlists

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example/bucket/app/helpers"
	"example/bucket/app/models/item"
	"example/bucket/app/models/wishlist"
)

func AddItem(ctx *gin.Context) {
	usr, _ := helpers.GetUser(ctx)
	wishlistId, err := strconv.Atoi(ctx.Param("wishlist_id"))
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body := helpers.AddItemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var wishlst wishlist.Wishlist

	result := wishlist.Wishlists(usr.ID).First(&wishlst, wishlistId)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var itm item.Item

	itm.WishlistId = uint(wishlistId)
	itm.Name = body.Name
	itm.Url = body.Url
	itm.Provider = body.Provider

	if result := item.DB().Create(&itm); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &itm)
}

func UpdateItem(ctx *gin.Context) {
	usr, _ := helpers.GetUser(ctx)
	wishlistId := ctx.Param("wishlist_id")
	itemId := ctx.Param("item_id")

	body := helpers.UpdateItemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var wishlst wishlist.Wishlist

	if result := wishlist.Wishlists(usr.ID).First(&wishlst, wishlistId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var itm item.Item

	if result := item.ActiveItems().Where("id = ? AND wishlist_id = ?", itemId, wishlistId).Find(&itm); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	itm.Description = body.Description
	itm.Priority = body.Priority

	item.DB().Save(&itm)

	ctx.JSON(http.StatusOK, &itm)

}

func RemoveItem(ctx *gin.Context) {
	usr, _ := helpers.GetUser(ctx)
	wishlistId := ctx.Param("wishlist_id")
	itemId := ctx.Param("item_id")

	var wishlst wishlist.Wishlist

	if result := wishlist.Wishlists(usr.ID).First(&wishlst, wishlistId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var itm item.Item

	if result := item.ActiveItems().Where("id = ? AND wishlist_id = ?", itemId, wishlistId).Find(&itm); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	itm.IsActive = false
	item.DB().Save(&itm)

	ctx.Status(http.StatusOK)
}
