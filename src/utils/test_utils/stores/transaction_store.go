package stores

import (
	"me-wallet/src/models"
)

var TransactionStore transactionStoreInterface

var transactions = map[uint]*models.Transaction{}

type transactionStore struct{}

type transactionStoreInterface interface {
	StoreTransaction(transaction *models.Transaction)
	GetTransactionsForUser(user models.User) []*models.Transaction
}

func init() {
	TransactionStore = &transactionStore{}
}

func (t *transactionStore) StoreTransaction(transaction *models.Transaction) {
	transactions[transaction.ID] = transaction
}

func (t *transactionStore) GetTransactionsForUser(user models.User) []*models.Transaction {
	var list []*models.Transaction
	for _, transaction := range transactions {
		if transaction.FromUserID == user.ID || transaction.ToUserID == user.ID {
			list = append(list, transaction)
		}
	}
	return list
}
