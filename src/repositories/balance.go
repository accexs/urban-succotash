package repositories

import (
	"me-wallet/src/models"
)

var BalanceRepository balanceRepoInterface

type balanceRepoInterface interface {
	GetForUser(user models.User) (*models.Balance, error)
	UpdateCurrentAmount(user models.User, amount float32) (*models.Balance, error)
}

type balanceRepo struct{}

func init() {
	BalanceRepository = &balanceRepo{}
}

func (b *balanceRepo) GetForUser(user models.User) (*models.Balance, error) {
	var balance models.Balance
	err := models.DB.Where("user_id = ?", user.ID).First(&balance).Error
	return &balance, err
}

func (b *balanceRepo) UpdateCurrentAmount(user models.User, amount float32) (*models.Balance, error) {
	balance, err := b.GetForUser(user)
	if err != nil {
		return nil, err
	}
	balance.CurrentAmount = balance.CurrentAmount + amount
	models.DB.Save(&balance)
	return balance, nil
}
