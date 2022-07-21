package handlers_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"me-wallet/src/handlers"
	"me-wallet/src/helpers"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
)

type balanceRepoMock struct{}

var funcGetForUser func(user models.User) (*models.Balance, error)

func (s *balanceRepoMock) GetForUser(user models.User) (*models.Balance, error) {
	return funcGetForUser(user)
}

func TestGetBalance(t *testing.T) {
	repositories.BalanceRepository = &balanceRepoMock{}

	funcGetForUser = func(user models.User) (*models.Balance, error) {
		return &models.Balance{
			BaseModel: models.BaseModel{
				ID: 1,
			},
			CurrentAmount: 0,
			UserID:        1,
		}, nil
	}

	request, _ := http.NewRequest("GET", "/banking/balance", nil)
	response := httptest.NewRecorder()
	c := helpers.CreateTestServer(request, response)
	helpers.MockCheckJwt(c)

	handlers.GetBalance(c)

	//r := helpers.CreateTestServer()
	//r.Use(helpers.MockCheckJwt)
	//r.GET("/banking/balance", handlers.GetBalance)
	//w := httptest.NewRecorder()
	//req, _ := http.NewRequest("GET", "/banking/balance", nil)
	//r.ServeHTTP(w, req)

	var balance models.Balance
	_ = json.Unmarshal(response.Body.Bytes(), &balance)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, float32(0), balance.CurrentAmount)
}

func TestGetTransactions(t *testing.T) {

}

func TestTransferMoney(t *testing.T) {

}
