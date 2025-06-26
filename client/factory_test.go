package client

import (
	"testing"
)

func TestNewBlockchainClient(t *testing.T) {
	_, err := NewBlockchainClient("ethereum")
	if err != nil {
		t.Errorf("NewBlockchainClient(ethereum) failed: %v", err)
	}
	_, err = NewBlockchainClient("tron")
	if err != nil {
		t.Errorf("NewBlockchainClient(tron) failed: %v", err)
	}
	_, err = NewBlockchainClient("unknown")
	if err == nil {
		t.Error("expected error for unknown blockchain, got nil")
	}
}
