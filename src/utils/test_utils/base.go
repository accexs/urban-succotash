package test_utils

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"net/http"
	"net/http/httptest"
	"os"
)

func CreateTestServer(request *http.Request, response *httptest.ResponseRecorder) (*gin.Context, *gin.Engine) {
	c, r := gin.CreateTestContext(response)
	err := os.Setenv("APP_ENV", "test")
	if err != nil {
		panic("Error setting env variable")
	}
	c.Request = request
	return c, r
}

func MockUserContext(user models.User, c *gin.Context) {
	c.Set("user", user)
}
