package user

import (
	"example/bucket/app/models"
	"example/bucket/app/models/wishlist"
)

type User struct {
	models.CommonModelFields
	UserName  string              `gorm:"unique; default:null" json:"user_name"`
	Email     string              `gorm:"unique" json:"email"`
	MobileNo  string              `gorm:"unique; default:null" json:"mobile_no"`
	Password  string              `json:"password"`
	Wishlists []wishlist.Wishlist `gorm:"ForeignKey:UserId" json:"wishlists"`
}
