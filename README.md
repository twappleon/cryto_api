# Blockchain SDK Go

A Go SDK for interacting with Ethereum and Tron blockchains. This SDK provides a unified interface for common blockchain operations such as wallet management, token transfers, and smart contract interactions.

## Installation

```bash
go get github.com/blockchain-sdk-go
```

## Quick Start

```go
package main

import (
	"context"
	"log"
	"math/big"

	"github.com/blockchain-sdk-go/client"
	"github.com/blockchain-sdk-go/types"
)

func main() {
	// Create a new Ethereum client
	ethClient, err := client.NewBlockchainClient(client.Ethereum)
	if err != nil {
		log.Fatal(err)
	}
	defer ethClient.Close()

	// Connect to an Ethereum node
	ctx := context.Background()
	err = ethClient.Connect(ctx, "https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
	if err != nil {
		log.Fatal(err)
	}

	// Generate a new wallet
	privateKey, address, err := ethClient.(types.WalletManager).GenerateNewWallet()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Generated new wallet: %s", address)

	// Get ETH balance
	balance, err := ethClient.(types.TokenManager).GetNativeBalance(ctx, address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ETH Balance: %s", balance.String())

	// Send ETH
	amount := new(big.Float).SetFloat64(0.1)
	txHash, err := ethClient.(types.TokenManager).SendNativeToken(ctx, privateKey, "0xRecipientAddress", amount)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transaction sent: %s", txHash)
}
```

## Features

- **Unified Interface**: Common interface for both Ethereum and Tron blockchains
- **Wallet Management**: Generate and manage wallets
- **Token Operations**: Send and receive native tokens and ERC20/TRC20 tokens
- **Smart Contract Interaction**: Deploy and interact with smart contracts
- **Event Subscription**: Subscribe to blockchain events

## Supported Operations

### Wallet Operations
- Generate new wallets
- Import wallets from private keys
- Sign transactions

### Token Operations
- Get native token balance (ETH/TRX)
- Send native tokens
- Get token balance (ERC20/TRC20)
- Send tokens

### Smart Contract Operations
- Deploy contracts
- Call contract functions
- Subscribe to contract events

## Security Considerations

- Private keys are never stored
- All transactions are signed offline
- Input validation for all parameters
- Secure connection handling

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 