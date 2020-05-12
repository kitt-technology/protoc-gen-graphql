package graphql

import (
	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"time"
)

type Mutation interface {
	GetName() string
	GetType() graphql.Object
	GetArgs() graphql.FieldConfigArgument
	GetSuccessEvent() *string
	GetFailureEvent() *string
}

type ProtoConfig struct {
	Mutations []*graphql.Field
	Queries   []*graphql.Field
}

type MutationResolver func(command proto.Message, success proto.Message) (proto.Message, error)

func GrpcConnection(name string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		name+".default.svc.cluster.local:50051",
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(30*time.Second),
	)

	if err != nil {
		panic(err)
	}

	return conn
}
