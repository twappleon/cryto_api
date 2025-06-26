package client

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/blockchain-sdk-go/api/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/golang/protobuf/proto"
)

// TronClient 實作 BlockchainClient, WalletManager, TokenManager, ContractManager

type TronClient struct {
	nodeURL string
	client  *client.GrpcClient
}

var _ types.BlockchainClient = (*TronClient)(nil)
var _ types.WalletManager = (*TronClient)(nil)
var _ types.TokenManager = (*TronClient)(nil)
var _ types.ContractManager = (*TronClient)(nil)

// Connect 實作 BlockchainClient 介面
func (t *TronClient) Connect(ctx context.Context, url string) error {
	cli := client.NewGrpcClient(url)
	err := cli.Start()
	if err != nil {
		return err
	}
	t.nodeURL = url
	t.client = cli
	return nil
}

// GenerateNewWallet 實作 WalletManager 介面
func (t *TronClient) GenerateNewWallet() (privateKey string, addr string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privBytes := crypto.FromECDSA(key)
	tronAddr := address.PubkeyToAddress(key.PublicKey).String()
	return hex.EncodeToString(privBytes), tronAddr, nil
}

// GetNativeBalance 实现 TokenManager
func (t *TronClient) GetNativeBalance(ctx context.Context, addr string) (*big.Float, error) {
	if t.client == nil {
		return nil, errors.New("Tron client not connected")
	}
	tronAddr, err := address.Base58ToAddress(addr)
	if err != nil {
		return nil, err
	}
	acc, err := t.client.GetAccount(tronAddr.String())
	if err != nil {
		return nil, err
	}
	balance := new(big.Float).SetInt(big.NewInt(acc.Balance))
	return balance, nil
}

// SendNativeToken 实现 TokenManager
func (t *TronClient) SendNativeToken(ctx context.Context, fromPrivateKey, toAddress string, amount *big.Float) (string, error) {
	if t.client == nil {
		return "", errors.New("Tron client not connected")
	}
	priv, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return "", err
	}
	fromAddr := address.PubkeyToAddress(priv.PublicKey).String()
	amt := big.NewInt(0)
	amount.Int(amt)
	txn, err := t.client.Transfer(fromAddr, toAddress, amt.Int64())
	if err != nil {
		return "", err
	}
	_, err = t.client.Broadcast(txn.Transaction)
	if err != nil {
		return "", err
	}
	txid := getTronTxID(txn.Transaction)
	return txid, nil
}

// DeployContract 实现 ContractManager
func (t *TronClient) DeployContract(ctx context.Context, privateKey, bytecode, abiJSON string, constructorArgs []interface{}) (string, error) {
	// gotron-sdk 暂无直接合约部署API，需用 TriggerSmartContract 创建合约
	return "", errors.New("Tron contract deployment not implemented in gotron-sdk, please use TronBox or TronGrid")
}

// CallContract 实现 ContractManager
func (t *TronClient) CallContract(ctx context.Context, contractAddress, abiJSON, method string, params []interface{}) (interface{}, error) {
	// gotron-sdk 支持 TriggerConstantContract 调用合约方法
	return nil, errors.New("Tron contract call not implemented in this demo, see gotron-sdk TriggerConstantContract")
}

// TRC20 余额查询
func (t *TronClient) GetTRC20Balance(ctx context.Context, contractAddress, walletAddress string) (*big.Int, error) {
	// 需用 TriggerConstantContract 调用 balanceOf
	return nil, errors.New("TRC20 balanceOf not implemented in this demo")
}

// TRC20 转账
func (t *TronClient) TransferTRC20(ctx context.Context, privateKey, contractAddress, toAddress string, amount *big.Int) (string, error) {
	// 需用 TriggerSmartContract 调用 transfer
	return "", errors.New("TRC20 transfer not implemented in this demo")
}

// Close 實作 BlockchainClient 介面
func (t *TronClient) Close() error {
	// gotron-sdk 没有 Close 方法，留空
	return nil
}

// getTronTxID 计算 Tron 交易哈希（TxID）
func getTronTxID(tx *core.Transaction) string {
	raw, _ := proto.Marshal(tx.GetRawData())
	hash := sha256.Sum256(raw)
	return hex.EncodeToString(hash[:])
}
