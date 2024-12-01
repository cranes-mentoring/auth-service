package repository

import "auth-service/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}
