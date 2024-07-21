package services

import (
	"enube-challenge/packages/logging"
	"go.uber.org/zap"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	logger    = logging.Log.With(zap.String("context", "jwt_service"))
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
		logger.Warn("Failed to sign JWT", zap.Error(err))
		return "", err
	}

	logger.Info("Signed token", zap.String("token", tokenString))
	return tokenString, nil
}

func (s *JWTService) Verify(tokenString string) (*Claims, error) {
	claims := &Claims{}
	logger.Info("Verifying token", zap.String("token", tokenString))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		logger.Error("Failed to parse token", zap.String("token", tokenString), zap.Error(err))
		return nil, err
	}
	if !token.Valid {
		logger.Error("Invalid token", zap.String("token", tokenString))
		return nil, err
	}

	return claims, nil
}
