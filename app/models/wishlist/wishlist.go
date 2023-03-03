package wishlist

import (
	"example/bucket/app/models"
	"example/bucket/app/models/item"
	"example/bucket/config/db"
	"fmt"

	"gorm.io/gorm"
)

type Wishlist struct {
	models.CommonModelFields
	UserId      uint        `gorm:"default:null" json:"user_id"`
	Name        string      `json:"url"`
	Type        string      `json:"type"`
	Description string      `json:"description"`
	Archived    bool        `gorm:"default:false" json:"archived"`
	Token       string      `json:"token"`
	Items       []item.Item `gorm:"ForeignKey:WishlistId" json:"items"`
}

func DB() *gorm.DB {
	return db.DB.Model(&Wishlist{})
}

func (wishlst Wishlist) String() string {
	return fmt.Sprintf("%s - %d", wishlst.Name, wishlst.UserId)
}

func UnarchivedWishlist() *gorm.DB {
	return DB().Preload("Items", "is_active = ?", true).Where("archived = ?", false)
}

func UserWishlists(userId uint) *gorm.DB {
	result := UnarchivedWishlist().Where("user_id = ?", userId)
	return result
}

func GuestWishlists(token string) *gorm.DB {
	result := UnarchivedWishlist().Where("token = ? AND user_id is null", token)
	return result
}

func AssignUserToWishlists(guestToken string, userId uint) {
	GuestWishlists(guestToken).Update("user_id", userId)
}
