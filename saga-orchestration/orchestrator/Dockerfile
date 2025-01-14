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
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/orchestrator main.go

# Use a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/orchestrator /root/orchestrator

# Set executable permissions
RUN chmod +x /root/orchestrator

# Expose API port
EXPOSE 8080

# Run the application
CMD ["/root/orchestrator"]
