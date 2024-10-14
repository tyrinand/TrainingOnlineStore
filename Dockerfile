FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
COPY .env ./
RUN go mod download

RUN go build -o product-service ./cmd/service/main.go

CMD ["./product-service"]



