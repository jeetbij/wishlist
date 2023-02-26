package models

import (
	"gorm.io/gorm"
)

type Item struct {
	CommonModelFields
	Name		string 		`json:"name"`
	Url			string 		`json:"url"`
	Provider	string 		`json:"provider"`
	Description	string 		`json:"description"`
	Priority	string 		`json:"priority"`
	IsActive	bool		`gorm:"default:true" json:"is_active"`
	WishlistID 	uint   		`json:"wishlist_id"`
}

func ActiveItems(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}
