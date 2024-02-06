# ビルド環境
FROM golang:1.21 AS builder

# アプリケーションのソースコードを含むディレクトリを指定
WORKDIR /app

# モジュールをダウンロード
COPY go.mod ./
RUN go mod download

# ソースコードをコピー
COPY *.go ./

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# 実行環境
FROM alpine:latest  

# セキュリティアップデート
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ビルド環境からバイナリをコピー
COPY --from=builder /app/myapp .

# 8080ポートを開放
EXPOSE 8080

# アプリケーションを実行
CMD ["./myapp"]
