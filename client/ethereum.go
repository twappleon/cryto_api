package client

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/blockchain-sdk-go/api/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumClient 實作 BlockchainClient, WalletManager, TokenManager, ContractManager

type EthereumClient struct {
	rpcURL string
	client *ethclient.Client
}

var _ types.BlockchainClient = (*EthereumClient)(nil)
var _ types.WalletManager = (*EthereumClient)(nil)
var _ types.TokenManager = (*EthereumClient)(nil)
var _ types.ContractManager = (*EthereumClient)(nil)

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
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	privBytes := crypto.FromECDSA(key)
	privateKey = hex.EncodeToString(privBytes)
	address = crypto.PubkeyToAddress(key.PublicKey).Hex()
	return privateKey, address, nil
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
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return ethValue, nil
}

// SendNativeToken 实现 TokenManager
func (e *EthereumClient) SendNativeToken(ctx context.Context, fromPrivateKey, toAddress string, amount *big.Float) (string, error) {
	if e.client == nil {
		return "", errors.New("Ethereum client not connected")
	}
	priv, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return "", err
	}
	fromAddr := crypto.PubkeyToAddress(priv.PublicKey)
	nonce, err := e.client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return "", err
	}
	valueWei := new(big.Int)
	amountWei := new(big.Float).Mul(amount, big.NewFloat(1e18))
	amountWei.Int(valueWei)
	gasLimit := uint64(21000)
	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	to := common.HexToAddress(toAddress)
	tx := ethtypes.NewTransaction(nonce, to, valueWei, gasLimit, gasPrice, nil)
	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return "", err
	}
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), priv)
	if err != nil {
		return "", err
	}
	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

// DeployContract 实现 ContractManager
func (e *EthereumClient) DeployContract(ctx context.Context, privateKey, bytecode, abiJSON string, constructorArgs []interface{}) (string, error) {
	if e.client == nil {
		return "", errors.New("Ethereum client not connected")
	}
	priv, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	fromAddr := crypto.PubkeyToAddress(priv.PublicKey)
	nonce, err := e.client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return "", err
	}
	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", err
	}
	bytecodeBytes, err := hex.DecodeString(strings.TrimPrefix(bytecode, "0x"))
	if err != nil {
		return "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(priv, big.NewInt(1)) // 1=mainnet, 可根据实际链ID调整
	if err != nil {
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice
	address, _, _, err := bind.DeployContract(auth, parsedABI, bytecodeBytes, e.client, constructorArgs...)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

// CallContract 实现 ContractManager
func (e *EthereumClient) CallContract(ctx context.Context, contractAddress, abiJSON, method string, params []interface{}) (interface{}, error) {
	if e.client == nil {
		return nil, errors.New("Ethereum client not connected")
	}
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	contract := common.HexToAddress(contractAddress)
	callData, err := parsedABI.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		To:   &contract,
		Data: callData,
	}
	output, err := e.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := parsedABI.UnpackIntoInterface(&result, method, output); err != nil {
		return nil, err
	}
	return result, nil
}

// ERC20 余额查询
func (e *EthereumClient) GetERC20Balance(ctx context.Context, contractAddress, walletAddress string) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	contract := common.HexToAddress(contractAddress)
	data, err := parsedABI.Pack("balanceOf", common.HexToAddress(walletAddress))
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}
	output, err := e.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(output), nil
}

// ERC20 转账
func (e *EthereumClient) TransferERC20(ctx context.Context, privateKey, contractAddress, toAddress string, amount *big.Int) (string, error) {
	priv, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	fromAddr := crypto.PubkeyToAddress(priv.PublicKey)
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return "", err
	}
	data, err := parsedABI.Pack("transfer", common.HexToAddress(toAddress), amount)
	if err != nil {
		return "", err
	}
	nonce, err := e.client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return "", err
	}
	gasPrice, err := e.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}
	contract := common.HexToAddress(contractAddress)
	tx := ethtypes.NewTransaction(nonce, contract, big.NewInt(0), uint64(60000), gasPrice, data)
	chainID, err := e.client.NetworkID(ctx)
	if err != nil {
		return "", err
	}
	signedTx, err := ethtypes.SignTx(tx, ethtypes.NewEIP155Signer(chainID), priv)
	if err != nil {
		return "", err
	}
	err = e.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

// ERC20 ABI 常量
const ERC20ABI = `[
  {"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},
  {"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"}
]`

// Close 實作 BlockchainClient 介面
func (e *EthereumClient) Close() error {
	if e.client != nil {
		e.client.Close()
	}
	return nil
}
