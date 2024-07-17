package users

import models "enube-challenge/packages/models/users"

type UserRepository interface {
	CreateUser(user *models.Users) error
	FindByEmail(email string) (*models.Users, error)
}
