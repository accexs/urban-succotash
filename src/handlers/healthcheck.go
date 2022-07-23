package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"net/http"
	"os"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} models.Health
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Health{Status: "ok", Environment: os.Getenv("APP_ENV")})
}
