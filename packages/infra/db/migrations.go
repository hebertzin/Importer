package db

import (
	"enube-challenge/packages/infra/db/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Users{}, &models.Supplier{})
	if err != nil {
		return err
	}
	return nil
}
