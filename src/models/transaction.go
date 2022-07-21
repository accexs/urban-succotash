package models

type Transaction struct {
	BaseModel
	Amount    float32 `json:"amount" gorm:"notNull"`
	Reference string  `json:"reference"`
	UserID    uint    `json:"userID" gorm:"notNull"`
	User      User    `json:"-"`
}
