# Build stage
FROM golang:1.27.2-alpine3.24 AS builder
WORKDIR /app

ENV GOPROXY=https://mirrors.cloud.tencent.com/go/|https://mirrors.aliyun.com/goproxy/|https://goproxy.cn|direct
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.20.1/migrate.linux-amd64.tar.gz | tar xvz


# Run stage
FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080

# CMD 会当作 ENTRYPOINT 的附加参数
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
