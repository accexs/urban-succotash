package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"me-wallet/src/models"
	"me-wallet/src/services"
	"net/http"
	"strconv"
)

func CheckJwt(c *gin.Context) {
	token, err := services.VerifyToken(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "unable to parse claims")
			return
		}
		var user models.User
		if err := models.DB.First(&user, userId); err != nil {
			c.Set("user", user)
		}
	}
	c.Next()
}
