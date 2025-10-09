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

# Build the graphql service
RUN go build -o /build/graphql-service ./cmd/graphql/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/graphql-service ./graphql

# Expose HTTP port
EXPOSE 8080

# Run the graphql service
CMD ["./graphql"]