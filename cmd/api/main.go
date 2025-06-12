package main

import (
	"log"

	"github.com/blockchain-sdk-go/api/handler"
	"github.com/blockchain-sdk-go/client"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Blockchain SDK API
// @version 1.0
// @description API for interacting with Ethereum and Tron blockchains
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Create Ethereum handler
	ethHandler, err := handler.NewBlockchainHandler(client.Ethereum)
	if err != nil {
		log.Fatalf("Failed to create Ethereum handler: %v", err)
	}

	// Create Tron handler
	tronHandler, err := handler.NewBlockchainHandler(client.Tron)
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
