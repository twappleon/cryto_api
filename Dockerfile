# 使用多階段建構，先編譯 Go 程式
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN cd cmd/api && go build -o /app/api-server main.go

# 使用 Nginx 作為反向代理，並將 Go 可執行檔複製進來
FROM nginx:alpine
WORKDIR /app
COPY --from=builder /app/api-server /app/api-server
COPY deploy/nginx.conf /etc/nginx/nginx.conf

# 開放 80 端口 (Nginx) 及 8080 (Go API)
EXPOSE 80 8080

# 啟動腳本：先啟動 Go API，再啟動 Nginx
CMD ["/bin/sh", "-c", "/app/api-server & nginx -g 'daemon off;'"] 