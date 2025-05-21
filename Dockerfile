FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o workout-app .

# Use a smaller image for the final build
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/workout-app .

# Copy .env file if needed
# COPY --from=builder /app/.env .

# Expose the port
EXPOSE 8000

# Run the application
CMD ["./workout-app"]