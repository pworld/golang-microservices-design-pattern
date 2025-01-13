FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o order-service main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order-service .

EXPOSE 8081

CMD ["./order-service"]
