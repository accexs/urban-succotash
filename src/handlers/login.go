package handlers

import (
	"github.com/gin-gonic/gin"
	"me-wallet/src/models"
	"me-wallet/src/services"
	"net/http"
)

// Login godoc
// @Summary Get access token.
// @Description Verifies credentials and returns access token.
// @Tags auth
// @Accept json
// @Produce json
// @Param Payload body models.Login true "Login"
// @Success 200 {object} models.TokenDetails
// @Failure 400 {object} object
// @Failure 422 {object} object
// @Router /login [post]
func Login(c *gin.Context) {
	var credentials models.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide valid login details"})
		return
	}

	if err := services.CheckPassword(credentials.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide valid login details"})
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
