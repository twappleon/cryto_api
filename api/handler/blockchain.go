package handler

import (
	"context"
	"math/big"
	"net/http"

	"github.com/blockchain-sdk-go/api/types"
	"github.com/blockchain-sdk-go/client"
	"github.com/gin-gonic/gin"
)

// BlockchainHandler 區塊鏈 API 處理器，負責處理 HTTP 請求
// client：區塊鏈操作介面
type BlockchainHandler struct {
	client  types.BlockchainClient
	nodeURL string
}

// NewBlockchainHandler 建立新的區塊鏈 handler，支援 nodeURL
// blockchainType：區塊鏈類型
// nodeURL：區塊鏈節點 URL
// 回傳 BlockchainHandler 實例與錯誤
func NewBlockchainHandler(blockchainType client.BlockchainType, nodeURL string) (*BlockchainHandler, error) {
	c, err := client.NewBlockchainClient(blockchainType)
	if err != nil {
		return nil, err
	}
	return &BlockchainHandler{client: c, nodeURL: nodeURL}, nil
}

// Connect 連接區塊鏈節點
// @Summary Connect to blockchain node
// @Description Connect to a blockchain node using the provided URL
// @Tags ethereum, tron
// @Accept json
// @Produce json
// @Param request body types.ConnectRequest true "Connection details"
// @Success 200 {object} types.Response
// @Router /eth/connect [post]
// @Router /tron/connect [post]
func (h *BlockchainHandler) Connect(c *gin.Context) {
	var req types.ConnectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	if err := h.client.Connect(c.Request.Context(), req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to connect",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		Code:    http.StatusOK,
		Message: "Connected successfully",
	})
}

// GenerateWallet 產生新錢包
// @Summary Generate new wallet
// @Description Generate a new blockchain wallet
// @Tags ethereum, tron
// @Produce json
// @Success 200 {object} types.Response{data=types.WalletResponse}
// @Router /eth/wallet/generate [post]
// @Router /tron/wallet/generate [post]
func (h *BlockchainHandler) GenerateWallet(c *gin.Context) {
	walletManager, ok := h.client.(types.WalletManager)
	if !ok {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Wallet operations not supported",
		})
		return
	}

	privateKey, address, err := walletManager.GenerateNewWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate wallet",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		Code:    http.StatusOK,
		Message: "Wallet generated successfully",
		Data: types.WalletResponse{
			Address:    address,
			PrivateKey: privateKey,
		},
	})
}

// GetBalance 查詢主鏈幣餘額
// @Summary Get native token balance
// @Description Get the balance of native tokens (ETH/TRX) for an address
// @Tags ethereum, tron
// @Accept json
// @Produce json
// @Param request body types.BalanceRequest true "Balance query details"
// @Success 200 {object} types.Response{data=types.BalanceResponse}
// @Router /eth/balance [post]
// @Router /tron/balance [post]
func (h *BlockchainHandler) GetBalance(c *gin.Context) {
	var req types.BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	tokenManager, ok := h.client.(types.TokenManager)
	if !ok {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Token operations not supported",
		})
		return
	}

	balance, err := tokenManager.GetNativeBalance(c.Request.Context(), req.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get balance",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	balanceFloat, _ := balance.Float64()
	c.JSON(http.StatusOK, types.Response{
		Code:    http.StatusOK,
		Message: "Balance retrieved successfully",
		Data: types.BalanceResponse{
			Address: req.Address,
			Balance: balanceFloat,
		},
	})
}

// SendNativeToken 發送主鏈幣
// @Summary Send native tokens
// @Description Send native tokens (ETH/TRX) to an address
// @Tags ethereum, tron
// @Accept json
// @Produce json
// @Param request body types.TransferRequest true "Transfer details"
// @Success 200 {object} types.Response{data=types.TransactionResponse}
// @Router /eth/transfer/native [post]
// @Router /tron/transfer/native [post]
func (h *BlockchainHandler) SendNativeToken(c *gin.Context) {
	var req types.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	tokenManager, ok := h.client.(types.TokenManager)
	if !ok {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Token operations not supported",
		})
		return
	}

	amount := new(big.Float).SetFloat64(req.Amount)
	txHash, err := tokenManager.SendNativeToken(c.Request.Context(), req.FromPrivateKey, req.ToAddress, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to send tokens",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		Code:    http.StatusOK,
		Message: "Transaction sent successfully",
		Data: types.TransactionResponse{
			TxHash: txHash,
		},
	})
}

// DeployContract 部署智能合約
// @Summary Deploy smart contract
// @Description Deploy a new smart contract to the blockchain
// @Tags ethereum, tron
// @Accept json
// @Produce json
// @Param request body types.ContractDeployRequest true "Contract deployment details"
// @Success 200 {object} types.Response{data=types.ContractResponse}
// @Router /eth/contract/deploy [post]
// @Router /tron/contract/deploy [post]
func (h *BlockchainHandler) DeployContract(c *gin.Context) {
	var req types.ContractDeployRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	contractManager, ok := h.client.(types.ContractManager)
	if !ok {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Contract operations not supported",
		})
		return
	}

	contractAddress, err := contractManager.DeployContract(
		c.Request.Context(),
		req.PrivateKey,
		req.Bytecode,
		req.ABI,
		req.ConstructorArgs,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to deploy contract",
			Data:    types.ErrorResponse{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		Code:    http.StatusOK,
		Message: "Contract deployed successfully",
		Data: types.ContractResponse{
			ContractAddress: contractAddress,
		},
	})
}

// ConnectByURL 供服务端自动连接节点
func (h *BlockchainHandler) ConnectByURL(url string) error {
	return h.client.Connect(context.Background(), url)
}

// Close 关闭区块链连接
func (h *BlockchainHandler) Close() {
	h.client.Close()
}
