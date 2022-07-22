package repositories

import (
	"me-wallet/src/models"
)

var UserRepository UserRepoInterface

type UserRepoInterface interface {
	GetById(id uint) (*models.User, error)
}

type userRepo struct{}

func init() {
	UserRepository = &userRepo{}
}

func (u *userRepo) GetById(id uint) (*models.User, error) {
	var user = models.User{}
	err := models.DB.First(&user, id).Error
	return &user, err
}
