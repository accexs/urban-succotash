package models

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	CurrencyType string `json:"currency" gorm:"notNull"`
	UserID       uint   `json:"user_id" gorm:"notNull"`
	User         User
}
