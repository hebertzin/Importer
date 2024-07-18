package users

import (
	"context"
	"enube-challenge/packages/errors"
	users_repository "enube-challenge/packages/interfaces/users"
	models "enube-challenge/packages/models/users"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	Create(ctx context.Context, user *models.Users) (*models.Users, error)
	FindByEmail(ctx context.Context, email string) (*models.Users, error)
}

type userService struct {
	repo users_repository.IUserRepository
}

func NewUsersService(repo users_repository.IUserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.Users) (*models.Users, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Some error has been ocurred", err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	err = s.repo.CreateUser(ctx, user)
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
