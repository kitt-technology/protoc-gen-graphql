# protoc-gen-graphql (PGG)

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
  "encoding/json"
  "fmt"
  "log"
  "net/http"

  "github.com/graphql-go/graphql"
  "github.com/yourproject/books"
  "google.golang.org/grpc"
)

type postData struct {
  Query     string                 `json:"query"`
  Operation string                 `json:"operation"`
  Variables map[string]interface{} `json:"variables"`
}

func main() {
  opts := []grpc.DialOption{grpc.WithInsecure()}

  // Initialize service and get GraphQL fields
  ctx, fields := books.Init(context.Background(), books.WithDialOptions(opts...))

  // Create GraphQL schema
  fieldMap := graphql.Fields{}
  for _, f := range fields {
    fieldMap[f.Name] = f
  }

  schema, err := graphql.NewSchema(graphql.SchemaConfig{
    Query: graphql.NewObject(graphql.ObjectConfig{
      Name:   "RootQuery",
      Fields: fieldMap,
    }),
  })
  if err != nil {
    log.Fatalf("failed to create schema: %v", err)
  }

  // Serve GraphQL endpoint
  http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
    var p postData
    if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
      w.WriteHeader(400)
      return
    }

    result := graphql.Do(graphql.Params{
      Context:        ctx,
      Schema:         schema,
      RequestString:  p.Query,
      VariableValues: p.Variables,
      OperationName:  p.Operation,
    })
    json.NewEncoder(w).Encode(result)
  })

  fmt.Println("GraphQL server running at http://localhost:8080/graphql")
  log.Fatal(http.ListenAndServe(":8080", nil))
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

Prevent N+1 queries by enabling batch loading. PGG supports two patterns:

#### Simple Batch Loading with `graphql.BatchRequest`

For simple cases where you just need to batch by string keys, use the built-in `graphql.BatchRequest`:

```protobuf
import "graphql.proto";

rpc LoadAuthors(graphql.BatchRequest) returns (AuthorsBatchResponse) {
  option (graphql.batch_loader) = true;
}

message AuthorsBatchResponse {
  map<string, Author> results = 1;  // Results keyed by request keys
}
```

The `graphql.BatchRequest` type contains a single field:
- `repeated string keys = 1` - The batch keys to load

#### Custom Batch Loading with Complex Request Types

For more complex batching scenarios where you need additional parameters:

```protobuf
rpc GetBooksBatch(GetBooksBatchRequest) returns (GetBooksBatchResponse) {
  option (graphql.batch_loader) = true;
}

message GetBooksBatchRequest {
  repeated GetBooksRequest reqs = 1;  // Custom request objects
}

message GetBooksBatchResponse {
  map<string, GetBooksResponse> results = 1;  // Results keyed by serialized request
}
```

The generated code automatically batches concurrent requests and deduplicates keys in both cases.

### Built-in Helper Types

PGG provides several built-in types in `graphql.proto` for common use cases:

#### `graphql.BatchRequest`

Use for simple batch loading by string keys:

```protobuf
message BatchRequest {
  repeated string keys = 1;
}
```

Example usage in proto:
```protobuf
rpc LoadAuthors(graphql.BatchRequest) returns (AuthorsBatchResponse) {
  option (graphql.batch_loader) = true;
}
```

**For service implementations**, import the proto-generated type:
```go
import (
    "github.com/kitt-technology/protoc-gen-graphql/graphql"
)

func (s *MyService) LoadAuthors(ctx context.Context, req *graphql.BatchRequest) (*AuthorsBatchResponse, error) {
    // req.Keys contains the batched keys
    // ...
}
```

**Note:** A runtime package (`github.com/kitt-technology/protoc-gen-graphql/runtime`) provides plain Go types for reference without importing proto files. However, gRPC service implementations must use the proto-generated types from the `graphql` package for proper serialization.

#### `graphql.PageInfo`

For pagination support:

```protobuf
message PageInfo {
  int32 total_count = 1;
  string end_cursor = 2;
  bool has_next_page = 3;
}
```

Example usage:
```protobuf
message UsersResponse {
  repeated User users = 1;
  graphql.PageInfo page_info = 2;
}
```

#### `graphql.FieldMask`

For optimizing queries with field selection:

```protobuf
message FieldMask {
  repeated string paths = 1;
  map<string, bool> paths_map = 2;
}
```

Example usage:
```protobuf
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
│(with DataLoader)│
└─────────┬───────┘
          │ gRPC
    ┌─────┴─────┬─────────┐
    ▼           ▼         ▼
┌────────┐ ┌────────┐ ┌────────┐
│Service1│ │Service2│ │Service3│
└────────┘ └────────┘ └────────┘
```

## 📚 Examples

See the [example](./example) directory for a comprehensive e-commerce example showcasing:

### Features Demonstrated

- **Multiple gRPC Services**: Products and Users microservices
- **Batch Loading with DataLoader**: Efficiently load related data without N+1 queries
- **Cross-Service Relationships**: Products linked to sellers (users)
- **Pagination Support**: Using `graphql.PageInfo` for paginated results
- **Field Masks**: Optimize queries with selective field loading
- **Complex Types**: Money, Inventory, Addresses, Timestamps, Enums
- **Optional Fields**: Properly handling nullable and optional values
- **Custom GraphQL Names**: Using `object_name` to rename types
- **Skipped Fields/Messages**: Internal data excluded from GraphQL schema

### Running with Docker Compose (Recommended)

The easiest way to run the example is using Docker Compose:

```bash
cd example
docker-compose up --build
```

This will start all three services:
- **Products service** (gRPC on port 50051) - Product catalog with inventory, pricing, and variants
- **Users service** (gRPC on port 50052) - Customer profiles, addresses, and loyalty info
- **GraphQL gateway** (HTTP on port 8080) - Unified GraphQL API

Once running, you'll see detailed test instructions in the console. Here are some example queries:

#### 1. Get All Users (Demonstrates Pagination)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ users_getUsers { users { id email firstName lastName type } pageInfo { totalCount hasNextPage } } }"}'
```

#### 2. Get Products with Inventory (Demonstrates Complex Types)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ products_getProducts { products { id name description category price { currencyCode units nanos } inventory { quantity reserved warehouseLocation } variants { name sku stockQuantity attributes } rating reviewCount } pageInfo { totalCount } } }"}'
```

#### 3. Search Products (Demonstrates Filtering & Pagination)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ products_searchProducts(query: \"wireless\", limit: 5) { products { id name price { units currencyCode } featured } pageInfo { hasNextPage endCursor } } }"}'
```

#### 4. Get User Profile (Demonstrates Nested Data)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ users_getUserProfile(userId: \"1\") { userId addresses { line1 city stateProvince postalCode country type isDefault } preferences { marketingEmails preferredLanguage preferredCurrency } loyalty { tier points discountPercentage } totalOrders memberSince } }"}'
```

#### 5. Products with Sellers (Demonstrates Cross-Service DataLoader)

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ products_getProducts { products { id name price { units currencyCode } seller { id email firstName lastName type } } } }"}'
```

To stop the services:

```bash
docker-compose down
```

### Running Locally

Alternatively, run the example locally:

```bash
make run-examples
# Visit http://localhost:8080/graphql
```

### Example Queries in GraphQL

```graphql
# Get products by category
query {
  products_getProducts(categories: [ELECTRONICS, SPORTS]) {
    products {
      id
      name
      category
      price {
        currencyCode
        units
        nanos
      }
      inventory {
        quantity
        warehouseLocation
      }
      seller {
        email
        firstName
        lastName
      }
    }
    pageInfo {
      totalCount
      hasNextPage
    }
  }
}

# Get user with full profile
query {
  users_getUserProfile(userId: "1") {
    userId
    addresses {
      line1
      city
      stateProvince
      country
      isDefault
    }
    loyalty {
      tier
      points
      discountPercentage
    }
    totalOrders
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

# Regenerate examples (uses buf)
make regenerate-examples

# Check if examples are up-to-date (useful in CI)
make check-examples
```

### Regenerating Example Code

After making changes to the code generator or templates, you need to regenerate the example code:

**Using Make (Recommended)**
```bash
make regenerate-examples
```

**Using Buf Directly**
```bash
go install .
buf generate --path example/
```

The CI pipeline automatically validates that examples are up-to-date. If you see a CI failure on the `check-examples` job, run `make regenerate-examples` and commit the changes.

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
- [ ] Performance benchmarks and optimizations

---

Made with ❤️ by [Kitt Technology](hstatttps://github.com/kitt-technology)