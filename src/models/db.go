package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=me_wallet_db port=5432 user=postgres password=postgres dbname=me_wallet sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = DB.AutoMigrate(&User{}, &Balance{})
	if err != nil {
		panic("Failed to run migrations")
	}

	DB = db
}
