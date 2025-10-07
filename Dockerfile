# Stage 1: Build the Go application
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application for a minimal final image
RUN CGO_ENABLED=0 GOOS=linux go build -o /coffee-chat-service ./main.go

# Stage 2: Create a minimal final image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /coffee-chat-service /coffee-chat-service

# Copy the public folder for static assets (uploaded images)
COPY --from=builder /app/public /public

# Expose port 8080 where the app runs
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/coffee-chat-service"]