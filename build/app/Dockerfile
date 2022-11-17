FROM golang:1.17.7-alpine
RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest
RUN mkdir -p /go/app
WORKDIR /go/app

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080
