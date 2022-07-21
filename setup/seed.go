package main

import (
	"me-wallet/src/models"
)

func main() {
	models.ConnectDatabase()

	user := models.User{
		Email:    "user@mail.com",
		Password: "password",
		Balance: []models.Balance{
			{CurrentAmount: 0},
		},
	}

	res := models.DB.Create(&user)
	if res.Error != nil {
		panic("Error seeding user")
	}
	println("Seeding completed")
}
