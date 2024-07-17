package users

import (
	"context"
	errors "enube-challenge/packages/errors"
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

func (s *userService) Create(ctx context.Context, user *models.Users) (*models.Users, error) {
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.ErrUserAlreadyExist
	}
	return user, nil
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*models.Users, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}