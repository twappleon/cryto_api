package client

import (
	"testing"
)

func TestErrors(t *testing.T) {
	if ErrUnsupportedBlockchain.Error() != "unsupported blockchain type" {
		t.Errorf("unexpected error string: %v", ErrUnsupportedBlockchain)
	}
}
