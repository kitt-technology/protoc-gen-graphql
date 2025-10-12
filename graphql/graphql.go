package graphql

import (
	"context"

	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

func GrpcConnection(host string, option ...grpc.DialOption) *grpc.ClientConn {
	// If no options are provided, use insecure credentials as default
	opts := option
	if len(opts) == 0 {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	conn, err := grpc.NewClient(
		host,
		opts...,
	)

	if err != nil {
		panic(err)
	}

	return conn
}

type GraphqlMessage interface {
	proto.Message
	XXX_GraphqlType() *gql.Object
	XXX_GraphqlArgs() gql.FieldConfigArgument
	XXX_Package() string
	FromArgs(args map[string]interface{})
}

// Module represents a proto package with services, messages, and types.
// Each generated proto file creates a Module that can be registered for GraphQL.
type Module interface {
	// Fields returns all GraphQL query/mutation fields from all services in this module
	Fields() gql.Fields

	// Messages returns all message types from this proto package
	Messages() []GraphqlMessage

	// WithLoaders registers all dataloaders from all services into the context
	WithLoaders(ctx context.Context) context.Context

	// PackageName returns the proto package name (e.g., "usersvc", "common")
	PackageName() string
}

// FieldCustomizer is a function that adds or modifies fields on a GraphQL type
type FieldCustomizer func(fields gql.FieldDefinitionMap)

// CombineFields merges fields from multiple modules into a single gql.Fields map.
// This is a helper function to easily combine fields from different proto packages.
//
// Example:
//
//	fields := pg.CombineFields(
//	    usersModule.Fields(),
//	    productsModule.Fields(),
//	    ordersModule.Fields(),
//	)
func CombineFields(fieldMaps ...gql.Fields) gql.Fields {
	combined := gql.Fields{}
	for _, fields := range fieldMaps {
		for name, field := range fields {
			combined[name] = field
		}
	}
	return combined
}

// CombineModuleFields is a convenience function that takes modules and combines their fields.
//
// Example:
//
//	fields := pg.CombineModuleFields(usersModule, productsModule, ordersModule)
func CombineModuleFields(modules ...Module) gql.Fields {
	fieldMaps := make([]gql.Fields, len(modules))
	for i, module := range modules {
		fieldMaps[i] = module.Fields()
	}
	return CombineFields(fieldMaps...)
}

// WithAllLoaders registers dataloaders from multiple modules into the context.
//
// Example:
//
//	ctx := pg.WithAllLoaders(ctx, usersModule, productsModule, ordersModule)
func WithAllLoaders(ctx context.Context, modules ...Module) context.Context {
	for _, module := range modules {
		ctx = module.WithLoaders(ctx)
	}
	return ctx
}

// DialOptions configures dial options for gRPC services.
// Service names should match the service names defined in your proto files.
// If no options are provided for a service, insecure credentials are used by default.
//
// Example:
//
//	dialOpts := pg.DialOptions{
//	    "Users":    []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
//	    "Products": []grpc.DialOption{grpc.WithTransportCredentials(yourTLSCreds)},
//	}
type DialOptions map[string][]grpc.DialOption
