package graphql

import (
	"context"
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

type Dataloader struct {
	Output graphql.Output
	Fn     DataloaderFn
}
type DataloaderFn func(context context.Context, ids []string) (interface{}, error)
type RegisterDataloaderFn func(typeDef Dataloader)

type ProtoConfig struct {
	Mutations  []*graphql.Field
	Queries    []*graphql.Field
	Dataloader map[string]DataloaderFn
}

type Svc interface {
	AppendDataloaders(map[string]Dataloader) map[string]Dataloader
}

func GrpcConnection(host string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		host,
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(30*time.Second),
	)

	if err != nil {
		panic(err)
	}

	return conn
}
