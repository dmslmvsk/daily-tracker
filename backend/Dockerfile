FROM golang:1.26-alpine

RUN apk add --no-cache git build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o main ./cmd/api/main.go


EXPOSE 8080


CMD ["./main"]