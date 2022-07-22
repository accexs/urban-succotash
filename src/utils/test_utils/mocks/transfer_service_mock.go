package mocks

import (
	"github.com/stretchr/testify/mock"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
)

type TransferServiceMock struct {
	mock.Mock
}

func (t *TransferServiceMock) MakeTransfer(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error) {
	transaction, _ := repositories.TransactionRepository.Store(toUser, fromUser, amount, reference)
	var balanceRepo = repositories.BalanceRepository
	_, _ = balanceRepo.UpdateCurrentAmount(toUser, amount)
	_, _ = balanceRepo.UpdateCurrentAmount(fromUser, -amount)
	return transaction, nil
}
