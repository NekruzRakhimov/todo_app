# Используем базовый образ Golang
FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todo_app

ENV PORT=8080

CMD ["./todo_app"]