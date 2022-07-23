package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	BaseModel
	Amount     float32 `json:"amount" gorm:"notNull" example:"250"`
	Reference  string  `json:"reference" example:"Transfer reference message"`
	FromUser   User    `json:"-"`
	FromUserID uint    `json:"fromUserID" gorm:"notNull" example:"123"`
	ToUser     User    `json:"-"`
	ToUserID   uint    `json:"toUserID" gorm:"notNull" example:"458"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	return nil
}

func (t *Transaction) AfterCreate(tx *gorm.DB) error {
	return nil
}
