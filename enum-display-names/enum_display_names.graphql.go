package enum_display_names

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*gql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*gql.Field
