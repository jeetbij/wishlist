package item

import (
	"example/bucket/app/models"

	"gorm.io/gorm"
)

type Item struct {
	models.CommonModelFields
	Name        string `json:"name"`
	Url         string `json:"url"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	WishlistId  uint   `json:"wishlist_id"`
}

func ActiveItems(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}
