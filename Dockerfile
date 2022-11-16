# ベースとなるDockerイメージの指定
FROM golang:1.17-alpine as server-build

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# ModuleモードをON
ENV GO111MODULE=on
WORKDIR  /go/Are-you-free
COPY . .

EXPOSE 8000

# alpineパッケージのアップデート
RUN apk upgrade --update && \
    apk --no-cache add git

# Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
