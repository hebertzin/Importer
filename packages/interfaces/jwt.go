package interfaces

import "github.com/golang-jwt/jwt/v5"

type Token interface {
	SignIn(email string) (string, error)
	Verify(tokenString string) (*jwt.Claims, error)
}
