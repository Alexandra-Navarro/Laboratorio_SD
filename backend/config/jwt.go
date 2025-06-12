package config

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

var jwtKey = []byte("secreto-super-seguro")

func GenerateJWT(rut string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   rut,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
