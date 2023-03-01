package migration

import (
	"example/bucket/app/models/item"
	"example/bucket/app/models/user"
	"example/bucket/app/models/wishlist"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&wishlist.Wishlist{}, &item.Item{}, &user.User{})
}
