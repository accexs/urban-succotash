package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"me-wallet/src/services"
	"net/http"
)

func Login(c *gin.Context) {
	var credentials models.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user registered"})
		return
	}

	if err := services.CheckPassword(credentials.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := services.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	return
}
