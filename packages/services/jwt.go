package services

import (
	"enube-challenge/packages/logging"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
	"time"
)

var (
	SecretKey = []byte(os.Getenv("SECRET_JWT"))
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) SignIn(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		logging.Log.Error("Failed to sign JWT", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService) Verify(tokenString string) (*Claims, error) {
	claims := &Claims{}
	logging.Log.Info("Verifying token", zap.String("token", tokenString))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		logging.Log.Error("Failed to parse token", zap.String("token", tokenString), zap.Error(err))
		return nil, err
	}
	if !token.Valid {
		logging.Log.Error("Invalid token", zap.String("token", tokenString))
		return nil, err
	}

	return claims, nil
}
