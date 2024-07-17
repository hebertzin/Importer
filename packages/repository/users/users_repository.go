package users

import (
	users "enube-challenge/packages/interfaces/users"
	models "enube-challenge/packages/models/users"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.UserRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) CreateUser(user *models.Users) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByEmail(email string) (*models.Users, error) {
	var user models.Users
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
