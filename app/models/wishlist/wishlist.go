package wishlist

import (
	"example/bucket/app/models"
	"example/bucket/app/models/item"

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

func UnarchivedWishlist(db *gorm.DB) *gorm.DB {
	return db.Where("archived = ?", false)
}
