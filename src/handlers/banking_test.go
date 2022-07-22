package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"me-wallet/src/handlers"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
	"me-wallet/src/services"
	"me-wallet/src/utils/test_utils"
	"me-wallet/src/utils/test_utils/mocks"
	"me-wallet/src/utils/test_utils/stores"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransferMoneyIntegration(t *testing.T) {
	models.ConnectTestDatabase()
	expectedFromUser := test_utils.UserFactory()
	models.DB.Create(&expectedFromUser)
	expectedFromUserBalance := test_utils.BalanceFactory(*expectedFromUser, 100.5)
	models.DB.Create(&expectedFromUserBalance)
	expectedToUser := test_utils.UserFactory()
	models.DB.Create(&expectedToUser)
	expectedToUserBalance := test_utils.BalanceFactory(*expectedToUser, 100)
	models.DB.Create(&expectedToUserBalance)

	body := handlers.TransferRequest{
		ToUserID:  expectedToUser.ID,
		Amount:    50.5,
		Reference: "An integration test transfer",
	}
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPut, "/banking/send", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)

	test_utils.MockUserContext(*expectedFromUser, c)

	handlers.TransferMoney(c)

	var actualToUserBalance, actualFromUserBalance models.Balance
	models.DB.First(&actualToUserBalance, expectedToUserBalance.ID)
	models.DB.First(&actualFromUserBalance, expectedFromUserBalance.ID)

	var transaction models.Transaction
	_ = json.Unmarshal(response.Body.Bytes(), &transaction)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, 50.5, transaction.Amount)
	assert.Equal(t, expectedToUser.ID, transaction.ToUserID)
	assert.Equal(t, expectedFromUser.ID, transaction.FromUserID)
	assert.EqualValues(t, 150.5, actualToUserBalance.CurrentAmount)
	assert.EqualValues(t, 50, actualFromUserBalance.CurrentAmount)
}

func TestGetBalance(t *testing.T) {
	repositories.BalanceRepository = &mocks.BalanceRepoMock{}

	request, _ := http.NewRequest(http.MethodGet, "/banking/balance", nil)
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	expectedUser := test_utils.UserFactory()
	expectedBalance := test_utils.BalanceFactory(*expectedUser, 0)
	stores.BalanceStore.StoreBalance(expectedBalance)
	test_utils.MockUserContext(*expectedUser, c)

	handlers.GetBalance(c)

	var balance models.Balance
	_ = json.Unmarshal(response.Body.Bytes(), &balance)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, float32(0), balance.CurrentAmount)
	assert.Equal(t, expectedBalance.ID, balance.ID)
	assert.Equal(t, expectedUser.ID, balance.UserID)
}

func TestGetTransactions(t *testing.T) {
	repositories.TransactionRepository = &mocks.TransactionRepoMock{}
	request, _ := http.NewRequest(http.MethodGet, "/banking/transaction", nil)
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)
	expectedFromUser := *test_utils.UserFactory()
	expectedToUser := *test_utils.UserFactory()
	test_utils.MockUserContext(expectedFromUser, c)

	for i := 0; i < 2; i++ {
		transaction := test_utils.TransactionFactory(expectedToUser, expectedFromUser, 50)
		stores.TransactionStore.StoreTransaction(transaction)
	}

	handlers.GetTransactions(c)

	var transactions []models.Transaction
	_ = json.Unmarshal(response.Body.Bytes(), &transactions)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, float32(50), transactions[0].Amount)
	assert.Equal(t, expectedToUser.ID, transactions[0].ToUserID)
	assert.Equal(t, expectedFromUser.ID, transactions[0].FromUserID)
	assert.Equal(t, float32(50), transactions[1].Amount)
	assert.Equal(t, expectedFromUser.ID, transactions[1].FromUserID)
	assert.Equal(t, expectedFromUser.ID, transactions[0].FromUserID)
}

func TestTransferMoney(t *testing.T) {
	repositories.TransactionRepository = &mocks.TransactionRepoMock{}
	repositories.UserRepository = &mocks.UserRepoMock{}
	repositories.BalanceRepository = &mocks.BalanceRepoMock{}
	services.TransferService = &mocks.TransferServiceMock{}

	expectedFromUser := test_utils.UserFactory()
	stores.UserStore.StoreUser(expectedFromUser)
	expectedFromUserBalance := test_utils.BalanceFactory(*expectedFromUser, 100.5)
	stores.BalanceStore.StoreBalance(expectedFromUserBalance)
	expectedToUser := test_utils.UserFactory()
	stores.UserStore.StoreUser(expectedToUser)
	expectedToUserBalance := test_utils.BalanceFactory(*expectedToUser, 100)
	stores.BalanceStore.StoreBalance(expectedToUserBalance)

	body := handlers.TransferRequest{
		ToUserID:  expectedToUser.ID,
		Amount:    50.5,
		Reference: "A test transfer",
	}
	jsonBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPut, "/banking/send", bytes.NewBuffer(jsonBody))
	response := httptest.NewRecorder()
	c, _ := test_utils.CreateTestServer(request, response)

	test_utils.MockUserContext(*expectedFromUser, c)

	handlers.TransferMoney(c)

	var transaction models.Transaction
	_ = json.Unmarshal(response.Body.Bytes(), &transaction)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.EqualValues(t, 50.5, transaction.Amount)
	assert.Equal(t, expectedToUser.ID, transaction.ToUserID)
	assert.Equal(t, expectedFromUser.ID, transaction.FromUserID)
	assert.EqualValues(t, 150.5, expectedToUserBalance.CurrentAmount)
	assert.EqualValues(t, 50, expectedFromUserBalance.CurrentAmount)
}
