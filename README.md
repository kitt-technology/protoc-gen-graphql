# protoc-gen-graphql (PGG)

[![CircleCI](https://circleci.com/gh/kitt-technology/protoc-gen-graphql.svg?style=svg)](https://circleci.com/gh/kitt-technology/protoc-gen-graphql)
[![Go Report Card](https://goreportcard.com/badge/github.com/kitt-technology/protoc-gen-graphql)](https://goreportcard.com/report/github.com/kitt-technology/protoc-gen-graphql)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

PGG is a protoc plugin that generates a performant GraphQL server to seamlessly knit together your gRPC services. It automatically creates GraphQL schemas, types, and resolvers from your Protocol Buffer definitions, enabling you to expose your gRPC services through a unified GraphQL API.

## ✨ Features

- **🚀 Automatic Schema Generation** - Generate complete GraphQL schemas from proto files
- **⚡ Performance Optimized** - Built-in DataLoader support prevents N+1 queries
- **🔌 Seamless gRPC Integration** - Direct integration with existing gRPC services
- **🎯 Type Safety** - Leverage proto types for compile-time safety
- **🔄 Batch Loading** - Automatic batching with `batch_loader` option
- **📝 Customizable** - Control GraphQL output with proto annotations
- **🌐 Buf Schema Registry Support** - Published to BSR for easy dependency management
- **🔧 Flexible** - Works with both `buf` and `protoc`

## 📦 Installation

### Using Go Install

Assuming you have `$GOBIN` on your path:

```bash
go install github.com/kitt-technology/protoc-gen-graphql@latest
```

### From Source

```bash
git clone https://github.com/kitt-technology/protoc-gen-graphql.git
cd protoc-gen-graphql
make build
```

## 🚀 Quick Start

### 1. Define Your Proto File

```protobuf
syntax = "proto3";

package books;

import "graphql.proto";

service Books {
  option (graphql.host) = "localhost:50051";

  rpc GetBooks(GetBooksRequest) returns (GetBooksResponse) {}

  rpc GetBooksByAuthor(GetBooksByAuthorRequest) returns (GetBooksByAuthorResponse) {
    option (graphql.batch_loader) = true;
  }
}

message Book {
  string id = 1;
  string name = 2;
  string author_id = 3;
}

message GetBooksRequest {
  repeated string ids = 1 [(graphql.optional) = true];
}

message GetBooksResponse {
  repeated Book books = 1;
}
```

### 2. Generate GraphQL Code

**With Buf (Recommended):**

```bash
buf generate
```

**With protoc:**

```bash
protoc \
  --proto_path . \
  -I ./graphql \
  --go_out=. \
  --go-grpc_out=. \
  --graphql_out="lang=go:." \
  books.proto
```

### 3. Create Your GraphQL Server

```go
package main

import (
  "context"
  "net/http"

  "github.com/graphql-go/graphql"
  "github.com/yourproject/books"
  "google.golang.org/grpc"
)

func main() {
  opts := []grpc.DialOption{grpc.WithInsecure()}

  // Initialize service and get GraphQL fields
  ctx, fields := books.Init(context.Background(), books.WithDialOptions(opts...))

  // Create GraphQL schema
  schema, _ := graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
      Name:   "Query",
      Fields: fieldsToMap(fields),
    }),
  })

  // Serve GraphQL endpoint
  http.HandleFunc("/graphql", graphqlHandler(ctx, schema))
  http.ListenAndServe(":8080", nil)
}
```

## 📖 Usage

### With Buf (Recommended)

The graphql proto definitions are published to the Buf Schema Registry, so you don't need to manually manage the import path.

**1. Add the dependency to your `buf.yaml`:**

```yaml
version: v2
modules:
  - path: .
deps:
  - buf.build/kitt-technology/graphql
```

**2. Run `buf dep update` to fetch dependencies**

**3. Update your proto imports:**

```protobuf
import "graphql.proto";
```

**4. Create a `buf.gen.yaml`:**

```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: .
    opt: paths=source_relative
  - local: protoc-gen-go-grpc
    out: .
    opt: paths=source_relative
  - local: protoc-gen-graphql
    out: .
    opt: lang=go
```

**5. Generate code:**

```bash
buf generate
```

### With protoc (Legacy)

```bash
protoc \
  --proto_path . \
  -I ./graphql \
  -I ${GOPATH}/src \
  --go_out=. \
  --go-grpc_out=. \
  --graphql_out="lang=go:." \
  ./path/to/your/file.proto
```

## 🎛️ Configuration Options

### Proto Annotations

#### Service Options

```protobuf
service MyService {
  option (graphql.host) = "localhost:50051";  // gRPC service host
}
```

#### Method Options

```protobuf
rpc BatchLoad(BatchRequest) returns (BatchResponse) {
  option (graphql.batch_loader) = true;  // Enable DataLoader batching
}
```

#### Message Options

```protobuf
message User {
  option (graphql.object_name) = "UserType";   // Custom GraphQL type name
  option (graphql.mutation) = true;             // Include in mutations
  option (graphql.skip_message) = true;         // Skip GraphQL generation
}
```

#### Field Options

```protobuf
message Request {
  string id = 1 [(graphql.optional) = true];   // Mark as optional in GraphQL
  string internal = 2 [(graphql.skip_field) = true];  // Skip in GraphQL
}
```

#### File Options

```protobuf
option (graphql.disabled) = true;    // Disable GraphQL generation for file
option (graphql.package) = "myapp";  // Set GraphQL package name
```

### Code Generation Options

- `lang=go` - Target language (currently only Go is supported)
- `module=github.com/your/module` - Go module path

## 🔥 Advanced Features

### DataLoader / Batch Loading

Prevent N+1 queries by enabling batch loading:

```protobuf
rpc LoadUsers(LoadUsersRequest) returns (LoadUsersResponse) {
  option (graphql.batch_loader) = true;
}

message LoadUsersRequest {
  repeated string keys = 1;  // Batch keys
}

message LoadUsersResponse {
  map<string, User> results = 1;  // Results keyed by request keys
}
```

The generated code automatically batches concurrent requests and deduplicates keys.

### Pagination Support

Use the built-in `PageInfo` type:

```protobuf
import "graphql.proto";

message UsersResponse {
  repeated User users = 1;
  graphql.PageInfo page_info = 2;
}
```

### Field Masking

Optimize queries with field masks:

```protobuf
import "graphql.proto";

message GetUserRequest {
  string id = 1;
  graphql.FieldMask field_mask = 2;
}
```

### Cross-Service Relationships

Stitch together multiple services:

```go
// Add custom fields to link services
books.BookGraphqlType.AddFieldConfig("author", &graphql.Field{
  Type: authors.AuthorGraphqlType,
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    book := p.Source.(*books.Book)
    return authors.LoadAuthors(p, book.AuthorId)
  },
})
```

## 🏗️ Architecture

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │ GraphQL Query
       ▼
┌─────────────────┐
│  GraphQL Server │  ← Generated by protoc-gen-graphql
│  (with DataLoader)│
└─────────┬───────┘
          │ gRPC
    ┌─────┴─────┬─────────┐
    ▼           ▼         ▼
┌────────┐ ┌────────┐ ┌────────┐
│Service1│ │Service2│ │Service3│
└────────┘ └────────┘ └────────┘
```

## 📚 Examples

See the [example](./example) directory for a complete working example with:
- Multiple gRPC services (Authors, Books)
- Batch loading with DataLoader
- Cross-service relationships
- Custom field resolvers

Run the example:

```bash
make run-examples
# Visit http://localhost:8080/graphql
```

Example query:

```graphql
query {
  books(ids: ["1", "2"]) {
    id
    name
    author {
      id
      name
    }
  }
}
```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

### Development

```bash
# Install dependencies
make deps

# Run tests
make test

# Run linter
golangci-lint run

# Build examples
make build-examples
```

## 🔒 Security

Please report security vulnerabilities to our [security policy](.github/SECURITY.md).

## 📝 License

This project is licensed under the [MIT License](LICENSE).

## 🙏 Acknowledgments

Built with:
- [graphql-go/graphql](https://github.com/graphql-go/graphql)
- [graph-gophers/dataloader](https://github.com/graph-gophers/dataloader)
- [protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf)

## 📮 Support

- 🐛 [Report a bug](https://github.com/kitt-technology/protoc-gen-graphql/issues/new?template=bug_report.yml)
- 💡 [Request a feature](https://github.com/kitt-technology/protoc-gen-graphql/issues/new?template=feature_request.yml)
- 💬 [Start a discussion](https://github.com/kitt-technology/protoc-gen-graphql/discussions)

## 🗺️ Roadmap

- [ ] Support for more GraphQL features (subscriptions, unions, interfaces)
- [ ] TypeScript/JavaScript code generation
- [ ] Additional language support
- [ ] GraphQL Federation support
- [ ] Performance benchmarks and optimizations

---

Made with ❤️ by [Kitt Technology](https://github.com/kitt-technology)