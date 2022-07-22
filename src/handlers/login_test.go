package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"me-wallet/src/handlers"
	"me-wallet/src/models"
	"me-wallet/src/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	models.ConnectTestDatabase()
	user := test_utils.UserFactory()
	body := models.Login{
		Password: user.Password,
		Email:    user.Email,
	}
	jsonBody, _ := json.Marshal(body)
	models.DB.Create(&user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	handlers.Login(c)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestLoginWrongPassword(t *testing.T) {
	models.ConnectTestDatabase()
	user := test_utils.UserFactory()
	body := models.Login{
		Password: "dummy-password",
		Email:    user.Email,
	}
	jsonBody, _ := json.Marshal(body)
	models.DB.Create(&user)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	handlers.Login(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestLoginWrongEmail(t *testing.T) {
	models.ConnectTestDatabase()
	body := models.Login{
		Password: "some@email.com",
		Email:    "password",
	}
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	handlers.Login(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestLoginWrongParams(t *testing.T) {
	models.ConnectTestDatabase()
	type badLogin struct {
		Passwords string `json:"passwords"`
		Emails    string `json:"emails"`
	}
	body := badLogin{
		Passwords: "some@email.com",
		Emails:    "password",
	}
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	handlers.Login(c)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
