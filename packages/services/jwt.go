package services

import (
	"enube-challenge/packages/domains"
	"enube-challenge/packages/logging"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	SecretKey = []byte(os.Getenv("SECRET_JWT"))
)

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) SignIn(email string) (string, error) {
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

func (s *JWTService) Verify(tokenString string) (domains.Claims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		logging.Log.Error("Failed to parse token", zap.String("token", tokenString), zap.Error(err))
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
