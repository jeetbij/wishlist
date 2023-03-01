package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example/bucket/app/config"
)

var DB *gorm.DB

func Init() {
	config := config.GetConfig()
	url := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		config.Dialect,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
}
