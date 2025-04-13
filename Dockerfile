FROM golang:1.21-alpine AS builder

WORKDIR /app

# 依存関係をコピーしてインストール
COPY go.mod ./
# go.sumがある場合のみコピー
COPY go.sum* ./
RUN go mod download

# ソースコードをコピー
COPY *.go ./

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-mkr-getmetric .

# 最終的なイメージを作成
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ビルダーステージからバイナリをコピー
COPY --from=builder /app/go-mkr-getmetric .

# 実行コマンドを設定
ENTRYPOINT ["./go-mkr-getmetric"]
