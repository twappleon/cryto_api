package types

// Response 標準 API 回應結構
// Code：狀態碼
// Message：訊息
// Data：回傳資料
type Response struct {
	Code    int         `json:"code"`           // 狀態碼
	Message string      `json:"message"`        // 訊息
	Data    interface{} `json:"data,omitempty"` // 回傳資料
}

// WalletResponse 錢包操作回應結構
// Address：錢包地址
// PrivateKey：私鑰
type WalletResponse struct {
	Address    string `json:"address"`               // 錢包地址
	PrivateKey string `json:"private_key,omitempty"` // 私鑰
}

// BalanceResponse 餘額查詢回應結構
// Address：查詢地址
// Balance：餘額
type BalanceResponse struct {
	Address string  `json:"address"` // 查詢地址
	Balance float64 `json:"balance"` // 餘額
}

// TransactionResponse 交易回應結構
// TxHash：交易雜湊
type TransactionResponse struct {
	TxHash string `json:"tx_hash"` // 交易雜湊
}

// ContractResponse 智能合約操作回應結構
// ContractAddress：合約地址
// TxHash：交易雜湊
// Result：執行結果
type ContractResponse struct {
	ContractAddress string      `json:"contract_address"`  // 合約地址
	TxHash          string      `json:"tx_hash,omitempty"` // 交易雜湊
	Result          interface{} `json:"result,omitempty"`  // 執行結果
}

// ErrorResponse 錯誤回應結構
// Error：錯誤訊息
type ErrorResponse struct {
	Error string `json:"error"` // 錯誤訊息
}
