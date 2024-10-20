package services

import (
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/errors"
	"enube-challenge/packages/logging"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userService struct {
	repo domain.IUserRepository
}

func NewUsersService(repo domain.IUserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Log.Error("Some error has been ocurred", zap.Error(err))
		return nil, err
	}
	user.Password = string(hashedPassword)

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.ErrUserAlreadyExist
	}
	return user, nil
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}
