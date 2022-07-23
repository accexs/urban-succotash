package main

import (
	"me-wallet/src/models"
)

func main() {
	models.ConnectDatabase()

	user := models.User{
		Email:    "user1@mail.com",
		Password: "password",
		Balance: []models.Balance{
			{CurrentAmount: 100},
		},
	}

	user2 := models.User{
		Email:    "user2@mail.com",
		Password: "password",
		Balance: []models.Balance{
			{CurrentAmount: 100},
		},
	}

	res := models.DB.Create(&user)
	if res.Error != nil {
		panic("Error seeding user 1")
	}
	res2 := models.DB.Create(&user2)
	if res2.Error != nil {
		panic("Error seeding user 2")
	}
	println("Seeding completed")
}
