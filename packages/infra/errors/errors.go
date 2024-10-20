package errors

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrInvalidInput        = errors.New("invalid input")
	ErrFailedGenerateToken = errors.New("failed to generate token")
	ErrInvalidCredentials  = errors.New("invalid credentials")
)
