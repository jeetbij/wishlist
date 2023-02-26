package models

import (
	"gorm.io/gorm"
)

type Wishlist struct {
	CommonModelFields
	Name		string 	`json:"url"`
	Type		string 	`json:"type"`
	Description	string 	`json:"description"`
	Archived	bool	`gorm:"default:false" json:"archived"`
	Items    	[]Item 	`gorm:"ForeignKey:WishlistID" json:"items"`
}

func UnarchivedWishlist(db *gorm.DB) *gorm.DB {
	return db.Where("archived = ?", false)
}
