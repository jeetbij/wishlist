package item

import (
	"example/bucket/app/config/db"
	"example/bucket/app/models"
	"fmt"

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

func DB() *gorm.DB {
	return db.DB
}

func (itm Item) String() string {
	return fmt.Sprintf("%s - %s", itm.Name, itm.Url)
}

func ActiveItems() *gorm.DB {
	return DB().Where("is_active = ?", true)
}
