package types

import (
	"context"
	"math/big"
)

// BlockchainClient 定義區塊鏈連線操作的通用介面
// 提供連接與關閉節點的方法
// Connect 連接區塊鏈節點，參數為 context 與節點 URL，回傳錯誤
// Close 關閉與區塊鏈節點的連線
type BlockchainClient interface {
	// Connect 連接區塊鏈節點
	Connect(ctx context.Context, url string) error
	// Close 關閉連線
	Close()
}

// WalletManager 定義錢包相關操作介面
// 包含產生新錢包、匯入私鑰、簽名交易等功能
type WalletManager interface {
	// GenerateNewWallet 產生新錢包，回傳私鑰與地址
	GenerateNewWallet() (privateKey string, address string, err error)
	// LoadWalletFromPrivateKey 由私鑰匯入錢包，回傳地址
	LoadWalletFromPrivateKey(privateKey string) (address string, err error)
	// SignTransaction 離線簽名交易
	SignTransaction(rawTx []byte, privateKey string) (signedTx []byte, err error)
}

// TokenManager 定義代幣操作介面
// 包含查詢餘額、發送主鏈幣與代幣
type TokenManager interface {
	// GetNativeBalance 查詢主鏈幣餘額（ETH/TRX）
	GetNativeBalance(ctx context.Context, address string) (*big.Float, error)
	// SendNativeToken 發送主鏈幣（ETH/TRX）
	SendNativeToken(ctx context.Context, fromPrivateKey, toAddress string, amount *big.Float) (txHash string, err error)
	// GetTokenBalance 查詢代幣餘額（ERC20/TRC20）
	GetTokenBalance(ctx context.Context, address, contractAddress string) (*big.Float, error)
	// SendToken 發送代幣（ERC20/TRC20）
	SendToken(ctx context.Context, fromPrivateKey, toAddress, contractAddress string, amount *big.Float) (txHash string, err error)
}

// ContractManager 定義智能合約操作介面
// 包含部署、呼叫合約、訂閱事件
type ContractManager interface {
	// DeployContract 部署智能合約
	DeployContract(ctx context.Context, bytecode string, abi string, constructorArgs ...interface{}) (contractAddress string, txHash string, err error)
	// CallContractFunction 呼叫合約方法
	CallContractFunction(ctx context.Context, abi, contractAddress, method string, params []interface{}) ([]interface{}, error)
	// SubscribeToEvents 訂閱合約事件
	SubscribeToEvents(ctx context.Context, contractAddress string, eventName string) (<-chan interface{}, error)
}

// EthereumClient 以太坊專用操作介面，整合所有功能
type EthereumClient interface {
	BlockchainClient
	WalletManager
	TokenManager
	ContractManager
}

// TronClient 波場專用操作介面，整合所有功能，並額外支援廣播簽名交易
type TronClient interface {
	BlockchainClient
	WalletManager
	TokenManager
	ContractManager
	// BroadcastSignedTransaction 廣播已簽名交易
	BroadcastSignedTransaction(ctx context.Context, txData []byte) (txHash string, err error)
}

// Logger 接口定义日志功能
type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}
