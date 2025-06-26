package client

import (
	"context"
	"math/big"
	"testing"
)

func TestTronClient_Connect(t *testing.T) {
	client := &TronClient{}
	err := client.Connect(context.Background(), "grpc.trongrid.io:50051")
	if err != nil {
		t.Errorf("Connect failed: %v", err)
	}
}

func TestTronClient_GenerateNewWallet(t *testing.T) {
	client := &TronClient{}
	priv, addr, err := client.GenerateNewWallet()
	if err != nil {
		t.Errorf("GenerateNewWallet failed: %v", err)
	}
	if priv == "" || addr == "" {
		t.Error("private key or address is empty")
	}
}

func TestTronClient_GetNativeBalance(t *testing.T) {
	client := &TronClient{}
	// 需先 Connect
	_ = client.Connect(context.Background(), "grpc.trongrid.io:50051")
	_, err := client.GetNativeBalance(context.Background(), "TXYz7Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw")
	if err == nil {
		t.Error("expected error for invalid address, got nil")
	}
}

func TestTronClient_SendNativeToken(t *testing.T) {
	client := &TronClient{}
	_ = client.Connect(context.Background(), "grpc.trongrid.io:50051")
	_, err := client.SendNativeToken(context.Background(), "invalidprivkey", "TXYz7Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw2Qw", big.NewFloat(1))
	if err == nil {
		t.Error("expected error for invalid private key, got nil")
	}
}

func TestTronClient_Close(t *testing.T) {
	client := &TronClient{}
	if err := client.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}
}
