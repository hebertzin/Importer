package migrations

import (
	user "enube-challenge/packages/models/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&user.Users{})
	if err != nil {
		return err
	}
	return nil
}
