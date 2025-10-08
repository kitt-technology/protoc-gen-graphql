package graphql

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func GrpcConnection(host string, option ...grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.NewClient(
		host,
		option...,
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
