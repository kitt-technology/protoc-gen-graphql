# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /build

# Copy library code (needed for replace directive)
COPY go.mod go.sum ./
COPY graphql ./graphql

# Copy example code
COPY example/go.mod example/go.sum ./example/
WORKDIR /build/example
RUN go mod download

# Copy example source code
COPY example ./

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