FROM golang:1.16.12-alpine AS builder

ENV GOPROXY=https://goproxy.cn  \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download \
    && go mod tidy

COPY . .

RUN go build -ldflags="-w -s" -a -installsuffix cgo -o blog-apiserver ./cmd/api-server/

FROM scratch

ENV APISERVERPORT=8080
# 从 builder 镜像中二进制文件拷贝到目录中
COPY --from=builder /blog-apiserver /

EXPOSE $SERVERPORT

# 需要运行的命令
ENTRYPOINT ["/blog-apiserver"]