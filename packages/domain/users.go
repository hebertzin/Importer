package domain

import (
	"context"
	"enube-challenge/packages/models"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindByEmail(ctx context.Context, email string) (*models.Users, error)
}
