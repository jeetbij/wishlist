package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"example/bucket/config"
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
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}
}
