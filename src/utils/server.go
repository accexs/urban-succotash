package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	_ "me-wallet/docs"
	"me-wallet/src/handlers"
	"me-wallet/src/middlewares"
	"me-wallet/src/models"
	"os"
)

// CreateServer
// @title Me wallet API
// @version 1.0
// @description API doc and playground for me chat.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
