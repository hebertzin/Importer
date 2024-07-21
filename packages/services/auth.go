package services

import (
	"context"
	"enube-challenge/packages/domain"
	"enube-challenge/packages/logging"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService interface {
	Auth(ctx context.Context, email string, password string) (domain.HttpResponse, error)
}

type authService struct {
	repo       domain.IUserRepository
	jwtService *JWTService
}

func NewAuthService(repo domain.IUserRepository, jwtService *JWTService) *authService {
	return &authService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *authService) Auth(ctx context.Context, email string, password string) (domain.HttpResponse, error) {
	logger := logging.Log.With(zap.String("context", "auth_service"))
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return domain.HttpResponse{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.HttpResponse{
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		}, err
	}

	token, err := s.jwtService.SignIn(user.Email)
	if err != nil {
		return domain.HttpResponse{
			Message: "Failed to generate token",
			Code:    http.StatusInternalServerError,
			Body:    "",
		}, err
	}

	logger.Info("Token generated", zap.String("token", token))
	return domain.HttpResponse{
		Message: "Authentication successful",
		Code:    http.StatusOK,
		Body:    token,
	}, nil
}
