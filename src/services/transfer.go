package services

import (
	"me-wallet/src/models"
	"me-wallet/src/repositories"
)

var TransferService transferServiceInterface

type transferServiceInterface interface {
	MakeTransfer(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error)
}

type transferService struct{}

func init() {
	TransferService = &transferService{}
}

func (t *transferService) MakeTransfer(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error) {
	var tmpDB = models.DB
	models.DB = models.DB.Begin()
	transaction, err := repositories.TransactionRepository.Store(toUser, fromUser, amount, reference)
	if err != nil {
		models.DB.Rollback()
		return nil, err
	}
	var balanceRepo = repositories.BalanceRepository
	_, err = balanceRepo.UpdateCurrentAmount(toUser, amount)
	if err != nil {
		models.DB.Rollback()
		return nil, err
	}
	_, err = balanceRepo.UpdateCurrentAmount(fromUser, -amount)
	if err != nil {
		models.DB.Rollback()
		return nil, err
	}
	models.DB.Commit()
	models.DB = tmpDB
	return transaction, nil
}
