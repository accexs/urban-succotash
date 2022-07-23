package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=me_wallet_db port=5432 user=postgres password=postgres dbname=me_wallet sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	migrate(db)

	DB = db
}

func ConnectTestDatabase() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migrate(db)

	DB = db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Balance{}, &Transaction{})
	if err != nil {
		panic("Failed to run migrations")
	}
}

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey" example:"123"`
	CreatedAt time.Time      `json:"createdAt" example:"2022-07-22T13:07:41.24104Z"`
	UpdatedAt time.Time      `json:"updatedAt" example:"2022-07-22T13:07:41.24104Z"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
