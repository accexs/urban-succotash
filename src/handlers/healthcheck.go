package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"net/http"
	"os"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Health{Status: "ok", Environment: os.Getenv("APP_ENV")})
}
