# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy Go modules
COPY consumer-service/go.mod consumer-service/go.sum ./
RUN go mod tidy

# Copy the full application source code
COPY consumer-service/ .

# Build the Go binary
RUN go build -o event-processing-consumer .

# Final Stage
FROM alpine:latest
WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/event-processing-consumer /event-processing-consumer

# Ensure executable permissions
RUN chmod +x /event-processing-consumer

# Debug: Verify the file exists
RUN ls -lah /event-processing-consumer

# Expose port
EXPOSE 9090

# Run the binary
CMD ["/event-processing-consumer"]
