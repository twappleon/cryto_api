package client

import (
	"github.com/blockchain-sdk-go/api/types"
)

// BlockchainType 區塊鏈類型字串定義
type BlockchainType string

const (
	Ethereum BlockchainType = "ethereum" // 以太坊
	Tron     BlockchainType = "tron"     // 波場
)

// NewBlockchainClient 根據區塊鏈類型建立對應 client
// blockchainType：區塊鏈類型
// 回傳 BlockchainClient 與錯誤
func NewBlockchainClient(blockchainType BlockchainType) (types.BlockchainClient, error) {
	switch blockchainType {
	case Ethereum:
		return NewEthereumClient()
	case Tron:
		return NewTronClient()
	default:
		return nil, ErrUnsupportedBlockchain
	}
}

// NewEthereumClient 建立以太坊 client
func NewEthereumClient() (types.BlockchainClient, error) {
	return &EthereumClient{}, nil
}

// NewTronClient 建立波場 client
func NewTronClient() (types.BlockchainClient, error) {
	return &TronClient{}, nil
}
