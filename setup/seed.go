package setup

import (
	"me-wallet/src/models"
)

func _() {
	models.ConnectDatabase()

	res := models.DB.Create(&models.User{Username: "admin", Email: "admin@mail.com", Password: "password"})
	if res.Error != nil {
		panic("Error seeding user")
	}
}
