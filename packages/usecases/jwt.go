package usecases

import (
	"enube-challenge/packages/domains"
	"enube-challenge/packages/infra/logging"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var (
	SecretKey = []byte(os.Getenv("SECRET_JWT"))
)

type JWTUseCase struct{}

func NewJWTUseCase() *JWTUseCase {
	return &JWTUseCase{}
}

func (s *JWTUseCase) SignIn(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	})
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		logging.Log.Error("Failed to sign JWT", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (s *JWTUseCase) Verify(tokenString string) (domains.Claims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		logging.Log.Error("Failed to parse user token", zap.String("token", tokenString), zap.Error(err))
		return domains.Claims{}, err
	}
	if !token.Valid {
		logging.Log.Error("Invalid token", zap.String("token", tokenString))
		return domains.Claims{}, err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return domains.Claims{}, fmt.Errorf("invalid token claims")
	}
	return domains.Claims{Email: email}, nil
}
