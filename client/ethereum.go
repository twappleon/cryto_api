package client

import (
	"context"
	"errors"
	"math/big"

	"github.com/blockchain-sdk-go/api/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClient 實作 BlockchainClient, WalletManager, TokenManager
// 你可以根據實際需求擴充欄位

type EthereumClient struct {
	rpcURL string
	client *ethclient.Client
}

// 确保 EthereumClient 实现 BlockchainClient, WalletManager, TokenManager
var _ types.BlockchainClient = (*EthereumClient)(nil)
var _ types.WalletManager = (*EthereumClient)(nil)
var _ types.TokenManager = (*EthereumClient)(nil)

// Connect 實作 BlockchainClient 介面
func (e *EthereumClient) Connect(ctx context.Context, url string) error {
	cli, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return err
	}
	e.rpcURL = url
	e.client = cli
	return nil
}

// GenerateNewWallet 實作 WalletManager 介面
func (e *EthereumClient) GenerateNewWallet() (privateKey string, address string, err error) {
	// 这里建议用 go-ethereum 的 crypto 包生成新私钥
	// 这里只返回假数据，实际可用 crypto.GenerateKey()
	return "dummy_eth_private_key", "dummy_eth_address", nil
}

// GetNativeBalance 实现 TokenManager
func (e *EthereumClient) GetNativeBalance(ctx context.Context, address string) (*big.Float, error) {
	if e.client == nil {
		return nil, errors.New("Ethereum client not connected")
	}
	acc := common.HexToAddress(address)
	balance, err := e.client.BalanceAt(ctx, acc, nil)
	if err != nil {
		return nil, err
	}
	// ETH 单位转换：Wei -> Ether
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return ethValue, nil
}

// SendNativeToken 实现 TokenManager
func (e *EthereumClient) SendNativeToken(ctx context.Context, fromPrivateKey, toAddress string, amount *big.Float) (string, error) {
	// 这里建议用 go-ethereum 的 bind 包和 crypto 包实现真实转账
	// 这里只返回假数据，实际实现较复杂
	return "dummy_tx_hash", nil
}

// Close 實作 BlockchainClient 介面
func (e *EthereumClient) Close() error {
	if e.client != nil {
		return e.client.Close()
	}
	return nil
}
