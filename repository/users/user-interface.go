package users

import "gin-gorm-ex/models"

type UserRepositoryInterface interface {
	CreateUser(newUser models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}
