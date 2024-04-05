package users

import (
	"errors"
	"gin-gorm-ex/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(newUser models.User) (models.User, error) {

	err := u.db.Save(&newUser).Error
	if err != nil || newUser.ID == 0 {
		return newUser, errors.New("Failed Create")
	}

	return newUser, nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {

	getUser := models.User{}

	if err := u.db.Where("email=? ", email).Find(&getUser).Error; err != nil || len(email) == 0 {
		return getUser, errors.New("Failed Get Data")
	}

	return getUser, nil
}
