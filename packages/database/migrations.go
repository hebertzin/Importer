package database

import (
	"enube-challenge/packages/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Users{}, &models.Supplier{})
	if err != nil {
		return err
	}
	return nil
}
