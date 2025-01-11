# Use a minimal Go image as the build stage
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application files
COPY . .

# Build the Go binary
RUN go build -o consumer-service .

# Use a lightweight image for the final container
FROM alpine:latest AS runtime
WORKDIR /root/
COPY --from=builder /app/consumer-service .

# Expose the necessary port
EXPOSE 8081

# Run the application
CMD ["./consumer-service"]
