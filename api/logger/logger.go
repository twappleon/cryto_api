package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/blockchain-sdk-go/types"
	"github.com/grafana/loki/clients/pkg/promtail/client"
	"github.com/sirupsen/logrus"
)

// LokiLogger 实现 Logger 接口，支持发送日志到 Loki
type LokiLogger struct {
	logger *logrus.Logger
	loki   client.Client
}

// NewLokiLogger 创建新的 Loki 日志记录器
func NewLokiLogger() (*LokiLogger, error) {
	logger := logrus.New()

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	// 设置日志级别
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	// 配置 Loki 客户端
	lokiURL := os.Getenv("LOKI_URL")
	if lokiURL == "" {
		lokiURL = "http://loki:3100"
	}

	cfg := client.Config{
		URL:       lokiURL,
		BatchWait: 1 * time.Second,
		BatchSize: 1024 * 1024,
		BackoffConfig: client.BackoffConfig{
			MinBackoff: 1 * time.Second,
			MaxBackoff: 5 * time.Second,
			MaxRetries: 5,
		},
		ExternalLabels: map[string]string{
			"service": "blockchain-sdk-api",
			"version": "1.0.0",
		},
	}

	lokiClient, err := client.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Loki client: %w", err)
	}

	return &LokiLogger{
		logger: logger,
		loki:   lokiClient,
	}, nil
}

// sendToLoki 发送日志到 Loki
func (l *LokiLogger) sendToLoki(level string, message string, fields map[string]interface{}) {
	labels := map[string]string{
		"level":   level,
		"service": "blockchain-sdk-api",
	}

	// 添加自定义标签
	for k, v := range fields {
		if str, ok := v.(string); ok {
			labels[k] = str
		}
	}

	entry := client.Entry{
		Labels: labels,
		Entry: client.Entry{
			Timestamp: time.Now(),
			Line:      message,
		},
	}

	// 异步发送到 Loki
	go func() {
		if err := l.loki.Send(context.Background(), entry); err != nil {
			l.logger.Errorf("Failed to send log to Loki: %v", err)
		}
	}()
}

// Info 记录信息级别日志
func (l *LokiLogger) Info(args ...interface{}) {
	message := fmt.Sprint(args...)
	l.logger.Info(message)
	l.sendToLoki("info", message, nil)
}

// Infof 记录格式化信息级别日志
func (l *LokiLogger) Infof(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Infof(format, args...)
	l.sendToLoki("info", message, nil)
}

// Error 记录错误级别日志
func (l *LokiLogger) Error(args ...interface{}) {
	message := fmt.Sprint(args...)
	l.logger.Error(message)
	l.sendToLoki("error", message, nil)
}

// Errorf 记录格式化错误级别日志
func (l *LokiLogger) Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Errorf(format, args...)
	l.sendToLoki("error", message, nil)
}

// Warn 记录警告级别日志
func (l *LokiLogger) Warn(args ...interface{}) {
	message := fmt.Sprint(args...)
	l.logger.Warn(message)
	l.sendToLoki("warn", message, nil)
}

// Warnf 记录格式化警告级别日志
func (l *LokiLogger) Warnf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Warnf(format, args...)
	l.sendToLoki("warn", message, nil)
}

// Debug 记录调试级别日志
func (l *LokiLogger) Debug(args ...interface{}) {
	message := fmt.Sprint(args...)
	l.logger.Debug(message)
	l.sendToLoki("debug", message, nil)
}

// Debugf 记录格式化调试级别日志
func (l *LokiLogger) Debugf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	l.logger.Debugf(format, args...)
	l.sendToLoki("debug", message, nil)
}

// WithField 添加单个字段
func (l *LokiLogger) WithField(key string, value interface{}) types.Logger {
	fields := map[string]interface{}{key: value}
	return l.WithFields(fields)
}

// WithFields 添加多个字段
func (l *LokiLogger) WithFields(fields map[string]interface{}) types.Logger {
	// 创建一个新的 logger 实例，包含字段信息
	newLogger := &LokiLogger{
		logger: l.logger,
		loki:   l.loki,
	}

	// 修改 sendToLoki 方法以包含字段
	originalSendToLoki := newLogger.sendToLoki
	newLogger.sendToLoki = func(level string, message string, additionalFields map[string]interface{}) {
		// 合并字段
		mergedFields := make(map[string]interface{})
		for k, v := range fields {
			mergedFields[k] = v
		}
		for k, v := range additionalFields {
			mergedFields[k] = v
		}
		originalSendToLoki(level, message, mergedFields)
	}

	return newLogger
}

// Close 关闭 Loki 客户端
func (l *LokiLogger) Close() error {
	if l.loki != nil {
		return l.loki.Stop()
	}
	return nil
}
