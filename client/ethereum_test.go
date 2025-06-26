package client

import (
	"context"
	"math/big"
	"testing"
)

func TestEthereumClient_Connect(t *testing.T) {
	client := &EthereumClient{}
	err := client.Connect(context.Background(), "https://mainnet.infura.io/v3/your-api-key")
	if err != nil {
		t.Errorf("Connect failed: %v", err)
	}
}

func TestEthereumClient_GenerateNewWallet(t *testing.T) {
	client := &EthereumClient{}
	priv, addr, err := client.GenerateNewWallet()
	if err != nil {
		t.Errorf("GenerateNewWallet failed: %v", err)
	}
	if priv == "" || addr == "" {
		t.Error("private key or address is empty")
	}
}

func TestEthereumClient_GetNativeBalance(t *testing.T) {
	client := &EthereumClient{}
	// 需先 Connect
	_ = client.Connect(context.Background(), "https://mainnet.infura.io/v3/your-api-key")
	_, err := client.GetNativeBalance(context.Background(), "0x0000000000000000000000000000000000000000")
	if err == nil {
		t.Error("expected error for invalid address, got nil")
	}
}

func TestEthereumClient_SendNativeToken(t *testing.T) {
	client := &EthereumClient{}
	_ = client.Connect(context.Background(), "https://mainnet.infura.io/v3/your-api-key")
	_, err := client.SendNativeToken(context.Background(), "invalidprivkey", "0x0000000000000000000000000000000000000000", big.NewFloat(1))
	if err == nil {
		t.Error("expected error for invalid private key, got nil")
	}
}

func TestEthereumClient_Close(t *testing.T) {
	client := &EthereumClient{}
	if err := client.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}
}
