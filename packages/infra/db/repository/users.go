package repository

import (
	"context"
	"enube-challenge/packages/domains"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) domains.UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) CreateUser(ctx context.Context, user *domains.User) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) FindByEmail(ctx context.Context, email string) (*domains.User, error) {
	var user domains.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
