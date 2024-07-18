package users_repository

import (
	"context"
	models "enube-challenge/packages/models/users"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindByEmail(ctx context.Context, email string) (*models.Users, error)
}
