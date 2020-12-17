package graphql

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
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

var Timestamp_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TimestampInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"ISOString": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

const name = "2006-01-02T15:04:05"

var Timestamp_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Timestamp",
	Fields: graphql.Fields{
		"ISOString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamp.Timestamp).Seconds, 0).Format("2006-01-02T15:04:05"), nil
			},
		},
	},
})

func ToTimestamp(field interface{}) *timestamp.Timestamp {
	timeMap := field.(map[string]interface{})
	t, _ := time.Parse("2006-01-02T15:04:05", timeMap["ISOString"].(string))
	ts := timestamp.Timestamp{
		Seconds: t.Unix(),
	}
	return &ts
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

type GraphqlMessage interface {
	proto.Message
	XXX_type() *graphql.Object
	XXX_args() graphql.FieldConfigArgument
	From_args(args map[string]interface{})
}

type Svc interface {
	AppendDataloaders(map[string]Dataloader) map[string]Dataloader
}

func GrpcConnection(host string, option ...grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.Dial(
		host,
		option...
	)

	if err != nil {
		panic(err)
	}

	return conn
}
