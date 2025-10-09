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

# Build the products service
RUN go build -o /build/products-service ./cmd/products/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/products-service ./products

# Expose gRPC port
EXPOSE 50051

# Run the products service
CMD ["./products"]