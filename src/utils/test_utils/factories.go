package test_utils

import (
	"github.com/brianvoe/gofakeit/v6"
	"me-wallet/src/models"
	"sync"
)

type autoInc struct {
	sync.Mutex
	id uint
}

func (ai *autoInc) GenID() (id uint) {
	ai.Lock()
	defer ai.Unlock()
	ai.id++
	id = ai.id
	return id
}

var ai autoInc

func UserFactory() *models.User {
	return &models.User{
		BaseModel: models.BaseModel{
			ID: ai.GenID(),
		},
		Email:    gofakeit.Email(),
		Password: "password",
	}
}

func BalanceFactory(user models.User, amount float32) *models.Balance {
	return &models.Balance{
		BaseModel: models.BaseModel{
			ID: ai.GenID(),
		},
		CurrentAmount: amount,
		UserID:        user.ID,
	}
}

func TransactionFactory(toUser models.User, fromUser models.User, amount float32) *models.Transaction {
	return &models.Transaction{
		BaseModel: models.BaseModel{
			ID: ai.GenID(),
		},
		Amount:     amount,
		Reference:  gofakeit.Sentence(5),
		ToUserID:   toUser.ID,
		FromUserID: fromUser.ID,
	}
}
