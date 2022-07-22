package mocks

import (
	"github.com/stretchr/testify/mock"
	"me-wallet/src/models"
	"me-wallet/src/utils/test_utils/stores"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) GetById(id uint) (*models.User, error) {
	user, err := stores.UserStore.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
