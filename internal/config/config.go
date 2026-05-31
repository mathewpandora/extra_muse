package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POLZA_API_KEY string
	BASE_URL string
	DATABASE_URL string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		POLZA_API_KEY: os.Getenv("POLZA_API_KEY"),
		BASE_URL: os.Getenv("BASE_URL"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
}