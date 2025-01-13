# Use Golang base image
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o payment-service main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/payment-service .

EXPOSE 8082

CMD ["./payment-service"]
