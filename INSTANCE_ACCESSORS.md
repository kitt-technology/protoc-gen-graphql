# Service Instance Accessors

The generated modules include a unified interface and accessor method that allows you to call service methods without needing to know whether you're using a gRPC client or a direct service implementation.

## API

For each service in your proto file, the module generates:

### `{ServiceName}Instance` interface

A unified interface containing all the service methods. This interface is compatible with both gRPC clients and service implementations through adapter types.

### `Get{ServiceName}() {ServiceName}Instance`

Returns a unified instance that implements the `{ServiceName}Instance` interface.
- Returns an adapter wrapping the gRPC client if one is configured
- Otherwise returns an adapter wrapping the service implementation
- Returns `nil` if neither is configured

## Example Usage

### Basic Usage

```go
package main

import (
    "context"
    "github.com/your-org/example/users"
)

func main() {
    // Create a module (with either client or service)
    module := users.NewUsersModule(
        users.WithModuleUsersService(myServiceImpl),
    )

    // Get the unified instance - no need to know if it's client or service!
    usersInstance := module.GetUsers()
    if usersInstance != nil {
        // Call methods directly - works with both client and service
        ctx := context.Background()
        resp, err := usersInstance.GetUsers(ctx, &users.GetUsersRequest{
            Ids: []string{"user1", "user2"},
        })
        if err != nil {
            // handle error
        }
        // use resp
    }
}
```

### With gRPC Client

```go
package main

import (
    "context"
    "google.golang.org/grpc"
    "github.com/your-org/example/users"
    pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
)

func main() {
    // Create a module with gRPC client configuration
    module := users.NewUsersModule(
        users.WithDialOptions(pg.DialOptions{
            "Users": []grpc.DialOption{grpc.WithInsecure()},
        }),
    )

    // Trigger lazy client creation
    _ = module.Fields()

    // Get the unified instance
    usersInstance := module.GetUsers()

    // Call methods - same API whether client or service!
    ctx := context.Background()
    profile, err := usersInstance.GetUserProfile(ctx, &users.GetUserProfileRequest{
        UserId: "user123",
    })
}
```

### Type Information

```go
// The interface is generated with all service methods:
type UsersInstance interface {
    GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error)
    GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*UserProfile, error)
}

// Adapters are generated automatically - you don't need to create them:
type usersClientAdapter struct { client UsersClient }
type usersServerAdapter struct { server UsersServer }
```

## Use Cases

1. **Abstraction**: Write code that works with both gRPC clients and direct service implementations
2. **Testing**: Easily swap between mock implementations and real clients
3. **Flexibility**: Change from client to service (or vice versa) without changing calling code
4. **Simplicity**: No type assertions or conditional logic needed in your code

## Notes

- The interface only includes public RPC methods (not internal loader methods)
- The adapters are generated automatically and handle the translation between client and server interfaces
- Client takes precedence if both client and service are somehow configured
- The returned instance is safe to use across goroutines if the underlying client/service is
- For lazy-loaded clients, the instance will only be available after the client has been created (e.g., after Fields() is called)

## Complete Example

```go
package main

import (
    "context"
    "log"
    "github.com/your-org/example/users"
)

// Function that works with any UsersInstance implementation
func printUserInfo(ctx context.Context, instance users.UsersInstance, userId string) {
    profile, err := instance.GetUserProfile(ctx, &users.GetUserProfileRequest{
        UserId: userId,
    })
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("User: %s %s", profile.User.FirstName, profile.User.LastName)
}

func main() {
    ctx := context.Background()

    // Works with service implementation
    serviceModule := users.NewUsersModule(
        users.WithModuleUsersService(myServiceImpl),
    )
    printUserInfo(ctx, serviceModule.GetUsers(), "user123")

    // Works with gRPC client too!
    clientModule := users.NewUsersModule(
        users.WithModuleUsersClient(myClient),
    )
    printUserInfo(ctx, clientModule.GetUsers(), "user456")
}
```