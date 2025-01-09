# Use Go as the build stage
FROM golang:1.23-alpine as builder

WORKDIR /app

# Install git (needed for some Go dependencies)
RUN apk add --no-cache git

# Copy go.mod and go.sum first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Debugging: List files to ensure everything is copied correctly
RUN ls -lah /app

# Build the Go binary
RUN go build -o /app/api-gateway-service .

# Use a minimal base image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api-gateway-service .

CMD ["./api-gateway-service"]
