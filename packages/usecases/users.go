package usecases

import (
	"context"
	"enube-challenge/packages/domains"
	"enube-challenge/packages/infra/errors"
	"enube-challenge/packages/infra/logging"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type usersUseCase struct {
	repo domains.UsersRepository
}

func NewUserUseCase(repo domains.UsersRepository) *usersUseCase {
	return &usersUseCase{
		repo: repo,
	}
}

func (s *usersUseCase) Create(ctx context.Context, user *domains.User) (*domains.User, error) {
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

func (s *usersUseCase) FindByEmail(ctx context.Context, email string) (*domains.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}
