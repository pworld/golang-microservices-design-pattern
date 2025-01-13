# Use Golang base image
FROM golang:1.23 as builder

# Set working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the application
RUN go build -o saga-orchestrator main.go

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/saga-orchestrator .

# Expose port if necessary
EXPOSE 8080

# Run the binary
CMD ["./saga-orchestrator"]
