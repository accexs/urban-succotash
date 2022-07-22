package handlers_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"me-wallet/src/handlers"
	"me-wallet/src/models"
	"me-wallet/src/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/check", nil)
	c, _ := test_utils.CreateTestServer(request, response)

	handlers.HealthCheck(c)

	var health models.Health
	_ = json.Unmarshal(response.Body.Bytes(), &health)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "ok", health.Status)
	assert.Equal(t, "test", health.Environment)
}
