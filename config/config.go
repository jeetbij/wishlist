package config

import (
	"log"
	"os"
	"strconv"

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

func LoadEnvVariables() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to load env")
	}
	env := os.Getenv("MODE")
	if env == "" {
		env = "debug"
	}

	err = godotenv.Load(rootPath + "/app/config/envs/" + env + "/.env")

	if err != nil {
		log.Println(err)
		log.Fatal("Failed to load .env for environment " + env)
	}
}

func GetConfig() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to get database port")
	}

	return &DBConfig{
		Dialect:  "postgres",
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		Username: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Database: os.Getenv("DATABASE_NAME"),
	}
}
