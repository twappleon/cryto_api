package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/blockchain-sdk-go/api/handler"
	"github.com/blockchain-sdk-go/client"
	_ "github.com/blockchain-sdk-go/cmd/api/docs" // 匿名 import，註冊 swagger 文件
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Blockchain SDK API
// @version 1.0
// @description API for interacting with Ethereum and Tron blockchains
// @host localhost:80
// @BasePath /api/v1
func main() {
	// 載入 .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	ethNodeURL := os.Getenv("ETH_NODE_URL")
	tronNodeURL := os.Getenv("TRON_NODE_URL")

	// Create Ethereum handler
	ethHandler, err := handler.NewBlockchainHandler(client.Ethereum, ethNodeURL)
	if err != nil {
		log.Fatalf("Failed to create Ethereum handler: %v", err)
	}

	// Create Tron handler
	tronHandler, err := handler.NewBlockchainHandler(client.Tron, tronNodeURL)
	if err != nil {
		log.Fatalf("Failed to create Tron handler: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// API version group
	v1 := r.Group("/api/v1")
	{
		// Ethereum routes
		eth := v1.Group("/eth")
		{
			eth.POST("/connect", ethHandler.Connect)
			eth.POST("/wallet/generate", ethHandler.GenerateWallet)
			eth.POST("/balance", ethHandler.GetBalance)
			eth.POST("/transfer/native", ethHandler.SendNativeToken)
			eth.POST("/contract/deploy", ethHandler.DeployContract)
		}

		// Tron routes
		tron := v1.Group("/tron")
		{
			tron.POST("/connect", tronHandler.Connect)
			tron.POST("/wallet/generate", tronHandler.GenerateWallet)
			tron.POST("/balance", tronHandler.GetBalance)
			tron.POST("/transfer/native", tronHandler.SendNativeToken)
			tron.POST("/contract/deploy", tronHandler.DeployContract)
		}
	}

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
