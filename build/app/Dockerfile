FROM golang:1.17.7-alpine
RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
RUN mkdir -p /go/app
WORKDIR /go/app

CMD ["air", "-c", ".air.toml"]
