package database

import (
	"enube-challenge/packages/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func ConnectDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Port, cfg.Host)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection successfully established")

	return db
}
