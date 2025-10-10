# GraphQL Input Type Override

## Overview

The `input_type_name` option allows you to override the auto-generated GraphQL input type for a message while keeping the output type auto-generated. This is useful when:

- The GraphQL input type needs to differ significantly from the proto message structure
- You want to add custom validation or field transformations on input
- You need custom descriptions or default values for input fields
- The input schema has additional computed fields or custom logic

## Usage

### 1. Add the `input_type_name` option to your proto message

```proto
syntax = "proto3";

package example;

import "graphql.proto";

message SearchRequest {
  option (graphql.input_type_name) = "CustomSearchInputType";

  string query = 1;
  int32 limit = 2;
  int32 offset = 3;
}
```

### 2. Create a custom input type in a `.graphql.go` file

Create a file alongside your generated code (e.g., `search-custom.graphql.go`) and define:

```go
package example

import (
    gql "github.com/graphql-go/graphql"
)

// Custom input type with additional validation or custom fields
var CustomSearchInputType = gql.NewInputObject(gql.InputObjectConfig{
    Name: "SearchRequestInput",
    Fields: gql.InputObjectConfigFieldMap{
        "query": &gql.InputObjectFieldConfig{
            Type:        gql.NewNonNull(gql.String),
            Description: "Search query string (required)",
        },
        "limit": &gql.InputObjectFieldConfig{
            Type:         gql.Int,
            DefaultValue: 10,
            Description:  "Maximum number of results (default: 10)",
        },
        "offset": &gql.InputObjectFieldConfig{
            Type:         gql.Int,
            DefaultValue: 0,
            Description:  "Number of results to skip (default: 0)",
        },
    },
})

// You still need to provide a FromArgs function for the message
// This is auto-generated, but you can customize it if needed
```

### 3. Generate your code

Run `protoc` with the `protoc-gen-graphql` plugin:

```bash
protoc --graphql_out=. --go_out=. your_file.proto
```

## What Gets Generated

When you use `input_type_name`:

✅ **Generated:**
- `SearchRequestGraphqlType` - The output type (auto-generated from proto)
- `SearchRequestFromArgs()` - Helper function to convert args to proto message
- `XXX_GraphqlType()` - Method for getting the output type

❌ **Not Generated:**
- `SearchRequestGraphqlInputType` - You provide this as `CustomSearchInputType`

Instead, the generator creates:
```go
// Using custom input type: CustomSearchInputType
var SearchRequestGraphqlInputType = CustomSearchInputType
```

## Complete Example

### Proto Definition

```proto
syntax = "proto3";

package products;

import "graphql.proto";

message ProductFilter {
  option (graphql.input_type_name) = "CustomProductFilterInput";

  string category = 1;
  int32 min_price = 2;
  int32 max_price = 3;
  bool in_stock = 4;
}

message Product {
  string id = 1;
  string name = 2;
  string category = 3;
  int32 price = 4;
  bool in_stock = 5;
}
```

### Custom Input Type

```go
// product-filter-custom.graphql.go
package products

import (
    gql "github.com/graphql-go/graphql"
)

var CustomProductFilterInput = gql.NewInputObject(gql.InputObjectConfig{
    Name: "ProductFilterInput",
    Description: "Custom filter for product searches with validation",
    Fields: gql.InputObjectConfigFieldMap{
        "category": &gql.InputObjectFieldConfig{
            Type:        gql.String,
            Description: "Product category (e.g., 'electronics', 'clothing')",
        },
        "minPrice": &gql.InputObjectFieldConfig{
            Type:        gql.Int,
            Description: "Minimum price in cents",
        },
        "maxPrice": &gql.InputObjectFieldConfig{
            Type:        gql.Int,
            Description: "Maximum price in cents",
        },
        "inStock": &gql.InputObjectFieldConfig{
            Type:         gql.Boolean,
            DefaultValue: true,
            Description:  "Only show in-stock items (default: true)",
        },
    },
})
```

## Comparison with `skip_message`

### Using `skip_message` (old approach)
```proto
message Money {
  option (graphql.skip_message) = true;
  string currency_code = 1;
  int64 units = 2;
}
```

- ❌ Skips **both** input and output type generation
- ❌ Must manually define both `MoneyGraphqlType` and `MoneyGraphqlInputType`
- ❌ Must manually define `MoneyFromArgs()`

### Using `input_type_name` (new approach)
```proto
message Money {
  option (graphql.input_type_name) = "CustomMoneyInput";
  string currency_code = 1;
  int64 units = 2;
}
```

- ✅ Only overrides the input type
- ✅ Output type is auto-generated
- ✅ `FromArgs()` is auto-generated
- ✅ Less boilerplate code

## When to Use

Use `input_type_name` when:
- You want different field names in the GraphQL input (e.g., camelCase vs snake_case)
- You need custom descriptions or default values for input fields
- You want to add input validation logic
- The input schema should differ from the output schema

Use `skip_message` when:
- You need complete control over both input and output types
- The message represents a custom scalar type
- You want to add computed fields to the output type

## Notes

1. The custom input type name you reference must be defined before the generated code is compiled
2. The `FromArgs()` function is still auto-generated to help convert from GraphQL args to proto messages
3. You can customize the `FromArgs()` behavior by defining your own if needed
4. The GraphQL args for fields using this type will still use the auto-generated field input types