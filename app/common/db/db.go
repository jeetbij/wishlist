package db

import (
	"log"
	"fmt"

	"gorm.io/driver/postgres"
    "gorm.io/gorm"

	"example/bucket/app/models"
	"example/bucket/app/common/config"
)

func Init() *gorm.DB {
	config := config.GetConfig()
	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		config.Dialect,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Wishlist{}, &models.Item{})

	return db
}
