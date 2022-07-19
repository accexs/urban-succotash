package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;notNull"`
	Email    string `json:"email" gorm:"unique;notNull"`
	Password string `json:"password;notNull"`
	Balance  []Balance
}

func (u *User) BeforeCreate(*gorm.DB) error {
	hashed, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
