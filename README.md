# Blockchain SDK Go

A Go SDK for interacting with Ethereum and Tron blockchains. This SDK provides a unified interface for common blockchain operations such as wallet management, token transfers, and smart contract interactions.

## Installation

```bash
go get github.com/blockchain-sdk-go
```

## Quick Start

### 使用 Docker Compose 啟動服務

1. 建置並啟動服務（在專案根目錄下執行）：
   ```bash
   docker-compose up --build
   ```
   此命令會啟動 Go API 服務（blockchain-api）與 Nginx 反向代理（blockchain-nginx），並將 Nginx 對外開放 80 端口。

2. 外部訪問 API 時，請使用 Nginx 反向代理的位址，例如：
   - 產生新錢包：  
     POST http://<your_host>/api/v1/eth/wallet/generate  
     (或 /api/v1/tron/wallet/generate)
   - 查詢餘額：  
     POST http://<your_host>/api/v1/eth/balance  
     (或 /api/v1/tron/balance)  
     (請求體範例：{ "address": "0x..." })
   - 發送主鏈幣：  
     POST http://<your_host>/api/v1/eth/transfer/native  
     (或 /api/v1/tron/transfer/native)  
     (請求體範例：{ "from_private_key": "0x...", "to_address": "0x...", "amount": 0.1 })
   - 部署合約：  
     POST http://<your_host>/api/v1/eth/contract/deploy  
     (或 /api/v1/tron/contract/deploy)  
     (請求體範例：{ "bytecode": "0x...", "abi": "..." })

3. Swagger 文件  
   可透過 http://<your_host>/swagger/index.html 查看 API 文件與測試。

### 注意事項

- 請確保 docker 與 docker-compose 已安裝，並在專案根目錄下執行 docker-compose 命令。
- 若需自訂 Nginx 反向代理設定，請修改 deploy/nginx.conf 檔案。
- 若需調整 Go API 服務的環境變數或參數，請在 docker-compose.yml 中設定。

---

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

## Swag init

```bash
swag init
```

# Swagger 文件產生與常見問題排查

## 產生 Swagger 文件

1. 安裝 swag CLI 工具：
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```
   > 安裝後請確認 `$GOPATH/bin` 或 `$HOME/go/bin` 已加入 PATH。

2. 驗證安裝：
   ```bash
   swag --version
   ```
   > 若有顯示版本號，表示安裝成功。

3. 回到專案根目錄（有 main.go 的地方），執行：
   ```bash
   swag init
   ```
   > 這會自動產生 `docs` 目錄與 swagger.json 文件。

## 常見問題排查

- 若執行 `swag init` 無法執行，請檢查：
  1. 是否已安裝 swag CLI 並加入 PATH。
  2. 專案根目錄是否有 main.go 並包含正確的 Swagger 註解（如 `@title`、`@version` 等）。
  3. 專案是否為 Go module（有 go.mod 檔）。
  4. 若有錯誤訊息，請參考終端機輸出內容進行修正。

- 若 Swagger UI 顯示 `Failed to load API definition` 或 `Internal Server Error doc.json`，請檢查：
  1. Gin 路由是否有註冊 `/swagger/*any` handler。
  2. `docs/swagger.json` 是否存在。
  3. Nginx 或 Docker 路由設定是否正確代理到 API 服務。
  4. API 服務日誌有無 500 錯誤或找不到檔案的訊息。

---

如遇問題，請將錯誤訊息貼給開發者協助排查。

---

## Swagger API 無法顯示接口時的排查與修復

如果你遇到 Swagger UI 顯示 `No operations defined in spec!` 或接口列表為空，請依照以下步驟排查：

### 1. 確認目錄結構
- 入口文件 `main.go` 位於 `cmd/api/main.go`
- handler 實現與註解位於 `api/handler/`

### 2. 正確執行 swag init 命令
請務必在**專案根目錄**下執行：
```bash
swag init --generalInfo cmd/api/main.go --output cmd/api/docs --dir api,cmd/api
```
- `--generalInfo` 指定入口 main.go 的正確路徑
- `--output` 指定 swagger 文件輸出目錄
- `--dir` 指定要掃描註解的多個目錄（用逗號分隔，無空格）

### 3. 常見錯誤與解法
- **路徑重複**：如果出現 `cannot parse source files .../cmd/api/cmd/api/main.go: no such file or directory`，說明路徑寫重複了，請確認只寫一次 `cmd/api`。
- **無 Go 檔案警告**：`no Go files in .../api` 只要 `api/handler` 有 Go 檔案即可，這個警告可忽略。
- **必須在根目錄執行**：建議始終在專案根目錄下執行 swag 命令，避免路徑混亂。

### 4. 生成後重啟服務
swag 生成文件後，請重啟 API 服務並刷新 Swagger UI 頁面。

--- 