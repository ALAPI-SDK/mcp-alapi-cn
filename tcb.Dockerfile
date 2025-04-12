FROM golang:1.24-alpine AS builder

# config alpine mirror
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories

# Install git since it might be needed for go mod
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files first for caching
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o mcp-alapi-cn .

# Final stage
FROM node:20-alpine

# config npm and yarn registry
RUN npm config set registry https://mirrors.cloud.tencent.com/npm/ \
    && yarn config set registry https://mirrors.cloud.tencent.com/npm/

# config alpine mirror
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tencent.com/g' /etc/apk/repositories

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/mcp-alapi-cn ./

# 安装 @cloudbase/mcp-transformer 工具
RUN npm install -g @cloudbase/mcp-transformer@1.0.0-beta.10

# 固定暴露端口
EXPOSE 80

# 启动命令
# 使用 cloudbase-mcp-transformer 将 Stdio 转换为远程 MCP 服务
CMD cloudbase-mcp-transformer stdioToCloudrun --stdioCmd "./mcp-alapi-cn" --port 80
