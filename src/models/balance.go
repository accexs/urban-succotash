package models

type Balance struct {
	BaseModel
	CurrentAmount float32 `json:"currentAmount" gorm:"notNull"`
	UserID        uint    `json:"userID" gorm:"notNull"`
	User          User    `json:"-"`
}
