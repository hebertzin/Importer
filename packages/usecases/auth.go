package usecases

import (
	"context"
	"enube-challenge/packages/domains"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type authUsecase struct {
	repo       domains.UsersRepository
	jwtUseCase *JWTUseCase
}

func NewAuthUseCase(repo domains.UsersRepository, jwtUseCase *JWTUseCase) *authUsecase {
	return &authUsecase{
		repo:       repo,
		jwtUseCase: jwtUseCase,
	}
}

func (s *authUsecase) Auth(ctx context.Context, email string, password string) (domains.HttpResponse, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return domains.HttpResponse{
			Message: "User not found",
			Code:    http.StatusNotFound,
		}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domains.HttpResponse{
			Message: "Invalid credentials",
			Code:    http.StatusUnauthorized,
		}, err
	}

	token, err := s.jwtUseCase.SignIn(user.Email)
	if err != nil {
		return domains.HttpResponse{
			Message: "Failed to generate token",
			Code:    http.StatusInternalServerError,
			Body:    "",
		}, err
	}

	return domains.HttpResponse{
		Message: "Authentication successful",
		Code:    http.StatusOK,
		Body:    token,
	}, nil
}
