# 使用官方的Golang基础镜像
FROM golang:1.17 as builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum到工作目录
COPY go.mod go.sum ./

# 下载项目依赖
RUN go mod download

# 复制项目源代码到工作目录
COPY . .

# 编译项目
RUN CGO_ENABLED=0 GOOS=linux go build -o love-scroll-api cmd/main.go

# 使用基于Alpine的轻量级镜像作为运行环境
FROM alpine:latest

# 在Alpine镜像中安装ca-certificates，以便支持TLS
RUN apk --no-cache add ca-certificates

# 将项目的可执行文件复制到Alpine镜像中
COPY --from=builder /app/love-scroll-api /app/

# 设置工作目录
WORKDIR /app

# 设置容器监听的端口
EXPOSE 8080

# 启动项目
CMD ["./love-scroll-api"]
