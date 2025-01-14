FROM golang:1.23 AS builder  

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/synchronization main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/synchronization /root/synchronization

RUN chmod +x /root/synchronization

EXPOSE 8083

CMD ["/root/synchronization"]
