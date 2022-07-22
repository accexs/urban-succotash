package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"me-wallet/src/handlers"
	"me-wallet/src/middlewares"
	"me-wallet/src/models"
	"os"
)

func CreateServer() *gin.Engine {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	f, _ := os.Create("./storage/app.log")
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/health", handlers.HealthCheck)
	r.POST("/login", handlers.Login)

	authorized := r.Group("/banking")

	authorized.Use(middlewares.CheckJwt)
	{
		authorized.GET("/balance", handlers.GetBalance)
		authorized.GET("/transactions", handlers.GetTransactions)
		authorized.PUT("/send", handlers.TransferMoney)
	}

	return r
}
