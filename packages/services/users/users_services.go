package users

import (
	"context"
	r "enube-challenge/packages/interfaces/users"
	models "enube-challenge/packages/models/users"
)

type UsersService interface {
	Create(ctx context.Context, user *models.Users) (*models.Users, error)
	FindByEmail(ctx context.Context, email string) (*models.Users, error)
}

type userService struct {
	repo r.UserRepository
}

func NewUsersService(repo r.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *models.Users) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*models.Users, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	return user, err
}
