package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/blockchain-sdk-go/api/handler"
	"github.com/blockchain-sdk-go/api/logger"
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
	// 初始化日志系统
	loggerInstance, err := logger.NewLokiLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer loggerInstance.Close()

	// 載入 .env
	if err := godotenv.Load(); err != nil {
		loggerInstance.Info("No .env file found, using system environment variables")
	}

	ethNodeURL := os.Getenv("ETH_NODE_URL")
	tronNodeURL := os.Getenv("TRON_NODE_URL")

	loggerInstance.Info("Starting Blockchain SDK API service")

	// Create Ethereum handler
	ethHandler, err := handler.NewBlockchainHandler(client.Ethereum, ethNodeURL)
	if err != nil {
		loggerInstance.Errorf("Failed to create Ethereum handler: %v", err)
		log.Fatalf("Failed to create Ethereum handler: %v", err)
	}

	// 自动连接以太坊节点
	if ethNodeURL != "" {
		loggerInstance.Infof("Auto-connecting Ethereum node: %s", ethNodeURL)
		if err := ethHandler.ConnectByURL(ethNodeURL); err != nil {
			loggerInstance.Errorf("Failed to auto-connect Ethereum node: %v", err)
			log.Fatalf("Failed to auto-connect Ethereum node: %v", err)
		}
		loggerInstance.Info("Successfully connected to Ethereum node")
	}

	// Create Tron handler
	tronHandler, err := handler.NewBlockchainHandler(client.Tron, tronNodeURL)
	if err != nil {
		loggerInstance.Errorf("Failed to create Tron handler: %v", err)
		log.Fatalf("Failed to create Tron handler: %v", err)
	}

	// 自动连接波场节点
	if tronNodeURL != "" {
		loggerInstance.Infof("Auto-connecting Tron node: %s", tronNodeURL)
		if err := tronHandler.ConnectByURL(tronNodeURL); err != nil {
			loggerInstance.Errorf("Failed to auto-connect Tron node: %v", err)
			log.Fatalf("Failed to auto-connect Tron node: %v", err)
		}
		loggerInstance.Info("Successfully connected to Tron node")
	}

	// Initialize Gin router
	r := gin.Default()

	// 添加日志中间件
	r.Use(func(c *gin.Context) {
		// 记录请求开始
		loggerInstance.Infof("Request started: %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()

		// 记录请求结束
		loggerInstance.Infof("Request completed: %s %s - Status: %d",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status())
	})

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

	// 设置优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		loggerInstance.Info("Starting HTTP server on port 8080")
		if err := r.Run(":8080"); err != nil {
			loggerInstance.Errorf("Failed to start server: %v", err)
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待关闭信号
	<-quit
	loggerInstance.Info("Shutting down server...")

	// 关闭区块链连接
	if ethHandler != nil {
		ethHandler.Close()
		loggerInstance.Info("Ethereum connection closed")
	}
	if tronHandler != nil {
		tronHandler.Close()
		loggerInstance.Info("Tron connection closed")
	}

	loggerInstance.Info("Server shutdown complete")
}
