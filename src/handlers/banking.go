package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"me-wallet/src/repositories"
	"net/http"
)

func GetBalance(c *gin.Context) {
	var user = c.MustGet("user").(models.User)
	//var balance models.Balance
	//if err := models.DB.Where("user_id = ?", user.ID).Preload("User").First(&balance).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{})
	//	return
	//}
	balance, err := repositories.BalanceRepository.GetForUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, balance)
	return
}

func GetTransactions(c *gin.Context) {

}

func TransferMoney(c *gin.Context) {

}
