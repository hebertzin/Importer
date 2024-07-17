package users

import (
	"context"
	models "enube-challenge/packages/models/users"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindByEmail(ctx context.Context, email string) (*models.Users, error)
}
