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