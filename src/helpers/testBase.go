package helpers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"net/http"
	"net/http/httptest"
	"os"
)

func CreateTestServer(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	err := os.Setenv("APP_ENV", "test")
	if err != nil {
		panic("Error setting env variable")
	}
	c.Request = request
	return c
}

func MockCheckJwt(c *gin.Context) {
	user := models.User{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Email:    "user@mail.com",
		Password: "Password",
		//Balance: []models.Balance{
		//	{
		//		BaseModel: models.BaseModel{
		//			ID: 1,
		//		},
		//		CurrentAmount: 0,
		//		UserID:        1,
		//	},
		//},
	}
	c.Set("user", user)
}
