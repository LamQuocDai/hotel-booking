package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL  string
	DBName string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{DBURL: os.Getenv("DB_URL"), DBName: os.Getenv("DB_Name")}, nil
}
