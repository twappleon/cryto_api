package client

import (
	"context"

	"github.com/blockchain-sdk-go/api/types"
)

// TronClient 實作 WalletManager 介面
// 你可以根據實際需求擴充欄位

// 確保 TronClient 實作 BlockchainClient, WalletManager
var _ types.BlockchainClient = (*TronClient)(nil)
var _ types.WalletManager = (*TronClient)(nil)

type TronClient struct {
	// 可擴充欄位
}

// Connect 實作 BlockchainClient 介面
func (t *TronClient) Connect(ctx context.Context, url string) error {
	// 這裡寫連線邏輯，暫時回傳 nil
	return nil
}

// GenerateNewWallet 實作 WalletManager 介面
func (t *TronClient) GenerateNewWallet() (privateKey string, address string, err error) {
	// 這裡寫產生錢包邏輯，這裡先回傳假資料
	return "dummy_private_key", "dummy_address", nil
}

// Close 實作 BlockchainClient 介面
func (t *TronClient) Close() error {
	// 如有資源釋放邏輯可寫於此，暫時回傳 nil
	return nil
}

// EthereumClient 實作 BlockchainClient, WalletManager
// 你可以根據實際需求擴充欄位

type EthereumClient struct {
	// 可擴充欄位
}

// 确保 EthereumClient 实现 BlockchainClient, WalletManager
var _ types.BlockchainClient = (*EthereumClient)(nil)
var _ types.WalletManager = (*EthereumClient)(nil)

// Connect 實作 BlockchainClient 介面
func (e *EthereumClient) Connect(ctx context.Context, url string) error {
	// 這裡寫連線邏輯，暫時回傳 nil
	return nil
}

// GenerateNewWallet 實作 WalletManager 介面
func (e *EthereumClient) GenerateNewWallet() (privateKey string, address string, err error) {
	// 這裡寫產生錢包邏輯，這裡先回傳假資料
	return "dummy_eth_private_key", "dummy_eth_address", nil
}

// Close 實作 BlockchainClient 介面
func (e *EthereumClient) Close() error {
	// 如有資源釋放邏輯可寫於此，暫時回傳 nil
	return nil
}
