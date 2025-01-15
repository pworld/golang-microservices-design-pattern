# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy Go modules
COPY producer-service/go.mod producer-service/go.sum ./
RUN go mod tidy

# Copy the full application source code
COPY producer-service/ .

# Build the Go binary
RUN go build -o event-processing-producer .

# Final Stage
FROM alpine:latest
WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/event-processing-producer /event-processing-producer

# Ensure executable permissions
RUN chmod +x /event-processing-producer

# Debug: Verify the file exists
RUN ls -lah /event-processing-producer

# Expose port
EXPOSE 8080

# Run the binary
CMD ["/event-processing-producer"]
