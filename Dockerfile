# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make protobuf protobuf-dev

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.34

# Build the plugin
RUN go build -o protoc-gen-graphql .

# Test stage (optional, can be used in CI)
FROM builder AS test
RUN make test

# Runtime stage for CI/testing
FROM golang:1.23-alpine AS runtime

# Install runtime dependencies
RUN apk add --no-cache \
    git \
    make \
    protobuf \
    protobuf-dev

# Copy built binaries from builder
COPY --from=builder /go/bin/protoc-gen-go /usr/local/bin/
COPY --from=builder /go/bin/protoc-gen-go-grpc /usr/local/bin/
COPY --from=builder /build/protoc-gen-graphql /usr/local/bin/

# Set up GOPATH
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/bin:$PATH

WORKDIR /workspace

# Default command
CMD ["/bin/sh"]