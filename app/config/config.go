package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func GetConfig() *DBConfig {
	return &DBConfig{
		Dialect:  "postgres",
		Host:     "127.0.0.1",
		Port:     5432,
		Username: "root",
		Password: "jeet@postgresql",
		Database: "bucket",
	}
}

func LoadEnvVariables() {
	var projectDirName = "bucket"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + "/app/common/envs/.env")

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to load env")
	}
}
