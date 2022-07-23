package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
	"me-wallet/src/services"
	"net/http"
)

// GetBalance godoc
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Summary User balance.
// @Description get the balance of logged-in user.
// @Tags banking
// @Accept */*
// @Produce json
// @Success 200 {object} models.Balance
// @Router /banking/balance [get]
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

// GetTransactions godoc
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Summary User transactions.
// @Description get transactions of logged-in user.
// @Tags banking
// @Accept */*
// @Produce json
// @Success 200 {array} models.Transaction
// @Router /banking/transactions [get]
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
	ToUserID  uint    `json:"toUserID" binding:"required" example:"123"`
	Amount    float32 `json:"amount" binding:"required,gt=0" example:"50"`
	Reference string  `json:"reference" example:"Transfer message example"`
}

// TransferMoney godoc
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Summary Make a transfer.
// @Description Transfer an amount of money from logged-in user to a target user.
// @Tags banking
// @Accept json
// @Produce json
// @Param Payload body TransferRequest true "Transfer"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} object
// @Failure 422 {object} object
// @Router /banking/send [put]
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
