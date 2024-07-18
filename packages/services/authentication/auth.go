package authentication

import (
	"context"

	"enube-challenge/packages/errors"
	users_repository "enube-challenge/packages/interfaces/users"
	"enube-challenge/packages/services/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Auth(ctx context.Context, email string, password string) (string, error)
}

type authService struct {
	repo       users_repository.IUserRepository
	jwtService *jwt.JWTService
}

func NewAuthService(repo users_repository.IUserRepository, jwtService *jwt.JWTService) *authService {
	return &authService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *authService) Auth(ctx context.Context, email string, password string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := s.jwtService.SignIn(user.Email)
	if err != nil {
		return "", errors.ErrFailedGenerateToken
	}

	return token, nil
}
