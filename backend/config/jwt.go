package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtKey string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JwtKey = os.Getenv("JWT_KEY")
}
