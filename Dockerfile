# FROM golang:1.16

# WORKDIR /app

# # コンテナ内の作業ディレクトリにgo.modとgo.sumをコピーして依存関係をダウンロードする
# COPY go.mod go.sum ./
# RUN go mod download

# # ローカルの全てのファイルをコンテナの作業ディレクトリにコピーする
# COPY . .

# # ビルド
# RUN go build -o main .

# # アプリケーションの起動コマンド
# CMD ["./main"]

# ベースイメージとして公式のGoイメージを使用
FROM golang:1.16

# コンテナ内の作業ディレクトリを設定
WORKDIR /go/src/app

# ホストのファイルをコンテナにコピー
COPY . .

# 必要な依存関係をダウンロード
RUN go mod download

# アプリケーションのビルド
RUN go build -o main

# コンテナの実行時のデフォルトコマンドを設定
CMD ["./main"]
