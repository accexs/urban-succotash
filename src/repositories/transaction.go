package repositories

import "me-wallet/src/models"

var TransactionRepository transactionRepoInterface

type transactionRepoInterface interface {
	Store(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error)
	GetForUser(user models.User) ([]*models.Transaction, error)
}

type transactionRepo struct{}

func init() {
	TransactionRepository = &transactionRepo{}
}

func (t transactionRepo) Store(toUser models.User, fromUser models.User, amount float32, reference string) (*models.Transaction, error) {
	transaction := &models.Transaction{
		Amount:     amount,
		Reference:  reference,
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
	}
	err := models.DB.Create(&transaction).Error
	return transaction, err
}

func (t transactionRepo) GetForUser(user models.User) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := models.DB.Where("from_user_id = ?", user.ID).Or("to_user_id = ?", user.ID).Find(&transactions).Error
	return transactions, err
}
