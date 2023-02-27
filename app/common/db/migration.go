package db

import (
	"example/bucket/app/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Wishlist{}, &models.Item{}, &models.User{})
}
