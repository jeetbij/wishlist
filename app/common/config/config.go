package config

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
