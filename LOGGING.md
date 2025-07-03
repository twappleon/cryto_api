# 日志配置说明

本项目集成了 Loki 日志系统，支持将应用日志发送到 Loki 进行集中管理和查询。

## 环境变量配置

在 `.env` 文件中添加以下配置：

```bash
# Blockchain Node URLs
ETH_NODE_URL=https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
TRON_NODE_URL=grpc.trongrid.io:50051

# Logging Configuration
LOG_LEVEL=info
LOKI_URL=http://loki:3100

# API Configuration
API_PORT=8080
```

## 日志级别

支持以下日志级别：
- `debug`: 调试信息
- `info`: 一般信息
- `warn`: 警告信息
- `error`: 错误信息

## 服务架构

使用 Docker Compose 启动完整的日志监控栈：

```bash
docker-compose up -d
```

这将启动以下服务：
- `blockchain-api`: 主应用服务
- `loki`: 日志聚合服务
- `promtail`: 日志收集器
- `grafana`: 日志可视化界面

## 访问 Grafana

1. 打开浏览器访问 `http://localhost:3000`
2. 使用默认凭据登录：
   - 用户名: `admin`
   - 密码: `admin`

## 配置 Loki 数据源

在 Grafana 中配置 Loki 数据源：

1. 进入 Settings > Data Sources
2. 点击 "Add data source"
3. 选择 "Loki"
4. 配置 URL: `http://loki:3100`
5. 点击 "Save & Test"

## 查询日志

在 Grafana Explore 中可以使用 LogQL 查询日志：

```logql
# 查询所有区块链 API 服务的日志
{service="blockchain-sdk-api"}

# 查询错误日志
{service="blockchain-sdk-api", level="error"}

# 查询特定 API 端点的日志
{service="blockchain-sdk-api"} |= "wallet/generate"

# 查询以太坊相关日志
{service="blockchain-sdk-api"} |= "ethereum"
```

## 日志标签

系统会自动添加以下标签：
- `service`: 服务名称 (blockchain-sdk-api)
- `level`: 日志级别 (debug, info, warn, error)
- `version`: 服务版本

## 自定义日志

在代码中使用日志：

```go
import "github.com/blockchain-sdk-go/api/logger"

// 获取日志实例
loggerInstance, err := logger.NewLokiLogger()
if err != nil {
    log.Fatal(err)
}
defer loggerInstance.Close()

// 记录日志
loggerInstance.Info("Service started")
loggerInstance.Errorf("Connection failed: %v", err)

// 添加字段
loggerInstance.WithField("user_id", "123").Info("User action")
loggerInstance.WithFields(map[string]interface{}{
    "tx_hash": "0x...",
    "amount": "1.5",
}).Info("Transaction completed")
```

## 故障排除

### Loki 连接失败
1. 检查 `LOKI_URL` 环境变量是否正确
2. 确保 Loki 服务正在运行
3. 检查网络连接

### 日志不显示
1. 检查 Promtail 配置
2. 确认日志文件路径正确
3. 验证 Loki 数据源配置

### 性能问题
1. 调整日志级别，减少不必要的日志
2. 优化 LogQL 查询
3. 考虑增加 Loki 资源限制 