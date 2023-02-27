package wishlists

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example/bucket/app/handlers/wishlists/helpers"
	"example/bucket/app/models"
)

func (h handler) AddItem(ctx *gin.Context) {
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

	var wishlist models.Wishlist

	result := h.DB.Scopes(models.UnarchivedWishlist).First(&wishlist, wishlistId)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var item models.Item

	item.WishlistID = uint(wishlistId)
	item.Name = body.Name
	item.Url = body.Url
	item.Provider = body.Provider

	if result := h.DB.Create(&item); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &item)
}

func (h handler) UpdateItem(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlist_id")
	itemId := ctx.Param("item_id")

	body := helpers.UpdateItemRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var wishlist models.Wishlist

	if result := h.DB.Scopes(models.UnarchivedWishlist).First(&wishlist, wishlistId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var item models.Item

	if result := h.DB.Scopes(models.ActiveItems).First(&item, itemId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	item.Description = body.Description
	item.Priority = body.Priority

	h.DB.Save(&item)

	ctx.JSON(http.StatusOK, &item)

}

func (h handler) RemoveItem(ctx *gin.Context) {
	wishlistId := ctx.Param("wishlist_id")
	itemId := ctx.Param("item_id")

	var wishlist models.Wishlist

	if result := h.DB.Scopes(models.UnarchivedWishlist).First(&wishlist, wishlistId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var item models.Item

	if result := h.DB.Scopes(models.ActiveItems).First(&item, itemId); result.Error != nil {
		log.Println(result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	item.IsActive = false
	h.DB.Save(&item)

	ctx.Status(http.StatusOK)
}
