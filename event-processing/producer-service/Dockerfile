# Build Stage: Use an official Golang image to build the application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the binary
RUN go build -o producer-service .

# Final Stage: Use a minimal image for production
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/producer-service .

# Expose the service port
EXPOSE 8080

# Run the producer service
CMD ["./producer-service"]
