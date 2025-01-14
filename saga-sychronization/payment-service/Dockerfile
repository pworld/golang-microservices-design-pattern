# Use Golang base image
FROM golang:1.23 AS builder  

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application source code
COPY . .

# Build a static binary compatible with Alpine Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/payment-service main.go

# Use a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/payment-service /root/payment-service

# Set executable permissions
RUN chmod +x /root/payment-service

# Expose API port (adjust as needed)
EXPOSE 8081

# Run the application
CMD ["/root/payment-service"]
