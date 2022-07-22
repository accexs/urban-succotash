package stores

import (
	"errors"
	"me-wallet/src/models"
)

var UserStore userStoreInterface

var users = map[uint]*models.User{}

type userStore struct{}

type userStoreInterface interface {
	StoreUser(user *models.User)
	GetUser(id uint) (*models.User, error)
}

func init() {
	UserStore = &userStore{}
}

func (u *userStore) StoreUser(user *models.User) {
	users[user.ID] = user
}

func (u *userStore) GetUser(id uint) (*models.User, error) {
	if user := users[id]; user != nil {
		return user, nil
	}
	return nil, errors.New("user not found")
}
