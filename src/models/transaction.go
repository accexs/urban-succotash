package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	BaseModel
	Amount     float32 `json:"amount" gorm:"notNull"`
	Reference  string  `json:"reference"`
	FromUser   User    `json:"-"`
	FromUserID uint    `json:"fromUserID" gorm:"notNull"`
	ToUser     User    `json:"-"`
	ToUserID   uint    `json:"toUserID" gorm:"notNull"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	// get user balance for FromUserID
	// check if (balance - amount) < 0
	// 	true -> return error
	return nil
}

func (t *Transaction) AfterCreate(tx *gorm.DB) error {
	//var balanceRepo = repositories.BalanceRepository
	//_, err := balanceRepo.UpdateCurrentAmount(t.ToUser, t.Amount)
	//if err != nil {
	//	return err
	//}
	//_, err = balanceRepo.UpdateCurrentAmount(t.FromUser, -t.Amount)
	//if err != nil {
	//	return err
	//}
	return nil
}
