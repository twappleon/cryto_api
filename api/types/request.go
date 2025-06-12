package types

// ConnectRequest 連接區塊鏈節點的請求結構
// URL：節點連線位址
type ConnectRequest struct {
	URL string `json:"url" binding:"required"` // 節點 URL
}

// WalletRequest 錢包操作請求結構
// PrivateKey：私鑰字串
type WalletRequest struct {
	PrivateKey string `json:"private_key" binding:"required"` // 私鑰
}

// TransferRequest 主鏈幣轉帳請求結構
// FromPrivateKey：發送方私鑰
// ToAddress：接收方地址
// Amount：轉帳金額
type TransferRequest struct {
	FromPrivateKey string  `json:"from_private_key" binding:"required"` // 發送方私鑰
	ToAddress      string  `json:"to_address" binding:"required"`       // 接收方地址
	Amount         float64 `json:"amount" binding:"required"`           // 轉帳金額
}

// TokenTransferRequest 代幣轉帳請求結構
// ContractAddress：代幣合約地址
type TokenTransferRequest struct {
	TransferRequest
	ContractAddress string `json:"contract_address" binding:"required"` // 代幣合約地址
}

// ContractDeployRequest 智能合約部署請求結構
// Bytecode：合約 bytecode
// ABI：合約 ABI
// ConstructorArgs：建構子參數
type ContractDeployRequest struct {
	Bytecode        string        `json:"bytecode" binding:"required"` // 合約 bytecode
	ABI             string        `json:"abi" binding:"required"`      // 合約 ABI
	ConstructorArgs []interface{} `json:"constructor_args"`            // 建構子參數
}

// ContractCallRequest 智能合約方法呼叫請求結構
// ContractAddress：合約地址
// ABI：合約 ABI
// Method：方法名稱
// Params：方法參數
type ContractCallRequest struct {
	ContractAddress string        `json:"contract_address" binding:"required"` // 合約地址
	ABI             string        `json:"abi" binding:"required"`              // 合約 ABI
	Method          string        `json:"method" binding:"required"`           // 方法名稱
	Params          []interface{} `json:"params"`                              // 方法參數
}

// BalanceRequest 查詢餘額請求結構
// Address：查詢地址
type BalanceRequest struct {
	Address string `json:"address" binding:"required"` // 查詢地址
}

// TokenBalanceRequest 查詢代幣餘額請求結構
// ContractAddress：代幣合約地址
type TokenBalanceRequest struct {
	BalanceRequest
	ContractAddress string `json:"contract_address" binding:"required"` // 代幣合約地址
}
