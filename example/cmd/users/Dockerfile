# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /build

# Copy go mod files from example directory
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . ./

# Build the users service
RUN go build -o /build/users-service ./cmd/users/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/users-service ./users

# Expose gRPC port
EXPOSE 50052

# Run the users service
CMD ["./users"]