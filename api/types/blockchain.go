package types

import (
	"context"
	"math/big"
)

// BlockchainClient 定義區塊鏈客戶端的基本介面
type BlockchainClient interface {
	// Connect 連接到區塊鏈節點
	Connect(ctx context.Context, url string) error
}

// WalletManager 定義錢包管理相關操作
type WalletManager interface {
	// GenerateNewWallet 生成新的錢包
	GenerateNewWallet() (privateKey string, address string, err error)
}

// TokenManager 定義代幣相關操作
type TokenManager interface {
	// GetNativeBalance 獲取主鏈幣餘額
	GetNativeBalance(ctx context.Context, address string) (*big.Float, error)
	// SendNativeToken 發送主鏈幣
	SendNativeToken(ctx context.Context, fromPrivateKey, toAddress string, amount *big.Float) (string, error)
}

// ContractManager 定義智能合約相關操作
type ContractManager interface {
	// DeployContract 部署智能合約
	DeployContract(ctx context.Context, privateKey, bytecode, abi string, constructorArgs []interface{}) (string, error)
	// CallContract 調用智能合約方法
	CallContract(ctx context.Context, contractAddress, abi, method string, params []interface{}) (interface{}, error)
}
