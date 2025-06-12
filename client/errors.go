package client

import "errors"

// ErrUnsupportedBlockchain 不支援的區塊鏈類型錯誤
var (
	// ErrUnsupportedBlockchain is returned when an unsupported blockchain type is specified
	ErrUnsupportedBlockchain = errors.New("unsupported blockchain type")
)
