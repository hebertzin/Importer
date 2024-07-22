package repository

import (
	"context"
	"enube-challenge/packages/domain"
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

func (r *usersRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
