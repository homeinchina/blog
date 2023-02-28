FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go mod download && go build -o server main.go

CMD ./server