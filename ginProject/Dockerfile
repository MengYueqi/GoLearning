# 使用官方Go镜像作为构建环境
FROM golang:1.22.5 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 复制到工作目录
COPY go.mod .
COPY go.sum .

# 下载所有依赖
RUN go mod download

# 将当前目录的所有文件复制到工作目录中
COPY . .

# 构建Go应用程序，确保输出二进制文件位于当前目录，并禁用CGO
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# 使用较小的镜像来运行应用程序
FROM alpine:latest

# 设置时区（可选）
RUN apk --no-cache add ca-certificates tzdata

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制所有应用程序文件到当前镜像，包括配置文件
COPY --from=builder /app .

# 确保可执行文件有执行权限（防止权限问题）
RUN chmod +x ./main

# 暴露应用运行的端口（根据您的应用设置）
EXPOSE 8080

# 运行应用程序
CMD ["./main"]
