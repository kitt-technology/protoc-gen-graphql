package graphql

import (
	"context"
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

var Timestamp_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Timestamp",
	Fields: graphql.InputObjectConfigFieldMap{
		"ISOString": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

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

type GraphqlMessage interface {
	proto.Message
	XXX_type() *graphql.Object
	XXX_args() graphql.FieldConfigArgument
	From_args(args map[string]interface{})
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
