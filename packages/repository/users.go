package repository

import (
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/models"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) domain.IUserRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) CreateUser(ctx context.Context, user *models.Users) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user models.Users
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
