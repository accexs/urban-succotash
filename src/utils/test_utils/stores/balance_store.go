package stores

import (
	"errors"
	"me-wallet/src/models"
)

var BalanceStore balanceStoreInterface

var balances = map[uint]*models.Balance{}

type balanceStore struct{}

type balanceStoreInterface interface {
	StoreBalance(balance *models.Balance)
	GetBalanceForUser(user models.User) (*models.Balance, error)
}

func init() {
	BalanceStore = &balanceStore{}
}

func (b *balanceStore) StoreBalance(balance *models.Balance) {
	balances[balance.ID] = balance
}

func (b *balanceStore) GetBalanceForUser(user models.User) (*models.Balance, error) {
	for _, balance := range balances {
		if balance.UserID == user.ID {
			return balance, nil
		}
	}
	return nil, errors.New("balance not found")
}
