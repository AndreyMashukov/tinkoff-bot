FROM golang:1.13

WORKDIR /go/src/app

COPY . /go/src/app

RUN go get github.com/joho/godotenv \
    && go get github.com/gorilla/websocket

RUN go build writer.go

CMD ["./writer"]
