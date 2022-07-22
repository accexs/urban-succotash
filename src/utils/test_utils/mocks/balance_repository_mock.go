package mocks

import (
	"github.com/stretchr/testify/mock"
	"me-wallet/src/models"
	"me-wallet/src/utils/test_utils/stores"
)

type BalanceRepoMock struct {
	mock.Mock
}

func (b *BalanceRepoMock) GetForUser(user models.User) (*models.Balance, error) {
	balance, err := stores.BalanceStore.GetBalanceForUser(user)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (b *BalanceRepoMock) UpdateCurrentAmount(user models.User, amount float32) (*models.Balance, error) {
	balance, err := b.GetForUser(user)
	if err != nil {
		return nil, err
	}
	balance.CurrentAmount = balance.CurrentAmount + amount
	return balance, nil
}
