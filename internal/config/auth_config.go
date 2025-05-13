package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AuthConfig struct {
	JwtSecretKey []byte
}

func LoadAuthConfig() (*AuthConfig, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	return &AuthConfig{
		JwtSecretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}, nil
}
