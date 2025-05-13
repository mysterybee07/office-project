package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	Schema   string
}

func LoadConfig() (*DatabaseConfig, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}, nil
}

func (c *DatabaseConfig) GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Schema,
	)
}
