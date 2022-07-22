package mocks

import (
	"github.com/stretchr/testify/mock"
	"me-wallet/src/models"
	"me-wallet/src/utils/test_utils/stores"
)

type TransactionRepoMock struct {
	mock.Mock
}

func (t *TransactionRepoMock) Store(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		Amount:     amount,
		Reference:  reference,
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
	}
	stores.TransactionStore.StoreTransaction(transaction)
	return transaction, nil
}

func (t *TransactionRepoMock) GetForUser(user models.User) ([]*models.Transaction, error) {
	transactions := stores.TransactionStore.GetTransactionsForUser(user)
	return transactions, nil
}
