# go_scraper/Dockerfile
FROM golang:1.20

# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# Goモジュールの初期化
RUN go mod init example.com/go_project

# .envファイルのコピー
COPY .env /go/src/app

COPY . /go/src/app

COPY . .
RUN go build -o main .

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]