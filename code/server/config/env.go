package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// .env file load
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
