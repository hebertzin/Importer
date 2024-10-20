package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User     string
	Password string
	Database string
	Port     string
	Host     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &Config{
		User:     os.Getenv("USER_DATABASE"),
		Password: os.Getenv("USER_PASSWORD"),
		Database: os.Getenv("DATABASE"),
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
	}

	fmt.Printf("Loaded Config: %+v\n", config)

	if config.User == "" || config.Password == "" || config.Database == "" || config.Port == "" || config.Host == "" {
		log.Fatal("One or more required environment variables are not set")
	}

	return config
}
