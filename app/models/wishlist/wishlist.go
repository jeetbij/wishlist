package wishlist

import (
	"example/bucket/app/config/db"
	"example/bucket/app/models"
	"example/bucket/app/models/item"
	"fmt"

	"gorm.io/gorm"
)

type Wishlist struct {
	models.CommonModelFields
	UserId      uint        `json:"user_id"`
	Name        string      `json:"url"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Archived    bool        `gorm:"default:false" json:"archived"`
	Items       []item.Item `gorm:"ForeignKey:WishlistId" json:"items"`
}

func DB() *gorm.DB {
	return db.DB
}

func (wishlst Wishlist) String() string {
	return fmt.Sprintf("%s - %d", wishlst.Name, wishlst.UserId)
}

func UnarchivedWishlist() *gorm.DB {
	return DB().Preload("Items", "is_active = ?", true).Where("archived = ?", false)
}

func Wishlists(userId uint) *gorm.DB {
	var result *gorm.DB
	if userId != 0 {
		result = UnarchivedWishlist().Where("user_id = ?", userId)
	} else {
		result = UnarchivedWishlist()
	}
	return result
}
