package handlers_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"me-wallet/src/handlers"
	"me-wallet/src/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCheck(t *testing.T) {
	r := gin.New()
	err := os.Setenv("APP_ENV", "test")
	if err != nil {
		panic("Error setting env variable")
	}
	r.GET("/check", handlers.HealthCheck)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/check", nil)
	r.ServeHTTP(w, req)

	var health models.Health
	_ = json.Unmarshal(w.Body.Bytes(), &health)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", health.Status)
	assert.Equal(t, "test", health.Environment)
}
