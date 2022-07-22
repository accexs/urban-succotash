package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
	"me-wallet/src/services"
	"net/http"
)

func GetBalance(c *gin.Context) {
	var user = c.MustGet("user").(models.User)
	balance, err := repositories.BalanceRepository.GetForUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, balance)
	return
}

func GetTransactions(c *gin.Context) {
	var user = c.MustGet("user").(models.User)
	transactions, err := repositories.TransactionRepository.GetForUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, transactions)
	return
}

type TransferRequest struct {
	ToUserID  uint    `json:"toUserID" binding:"required"`
	Amount    float32 `json:"amount" binding:"required,gt=0"`
	Reference string  `json:"reference"`
}

func TransferMoney(c *gin.Context) {
	var fromUser = c.MustGet("user").(models.User)
	var transferInput TransferRequest
	if err := c.ShouldBindJSON(&transferInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toUser, err := repositories.UserRepository.GetById(transferInput.ToUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipient user not found"})
		return
	}

	transaction, err := services.TransferService.MakeTransfer(*toUser, fromUser, transferInput.Amount, transferInput.Reference)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to process transfer"})
		return
	}
	c.JSON(http.StatusCreated, transaction)
	return
}
