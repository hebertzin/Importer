package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("USER_DATABASE")
	password := os.Getenv("USER_PASSWORD")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if user == "" || password == "" || database == "" || port == "" || host == "" {
		log.Fatal("One or more required environment variables are not set")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable TimeZone=Asia/Shanghai",
		user, password, database, port, host)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established")

	return db
}
