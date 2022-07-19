package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
)

func getBalance(c *gin.Context) {
	models.DB.Find(models.Balance{})
}
