FROM golang:1.19-alpine as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .


CMD ["air"]
