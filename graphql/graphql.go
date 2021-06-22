package graphql

import (
	"context"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"

	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
)

var Timestamp_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TimestampInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"ISOString": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

const name = "2006-01-02T15:04:05"

var WrappedString = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "WrappedString",
	Description: "protobuf string wrapper",
	Serialize: func(value interface{}) interface{} {
		return value.(*wrapperspb.StringValue).GetValue()
	},
	ParseValue: func(value interface{}) interface{} {
		// value is of type string... expected.
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		// GetValue() is of type *wrapperspb.StringValue, why?
		return valueAST.GetValue()
	},
})

var PageInfo_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfo",
	Fields: graphql.Fields{
		"total_count": &graphql.Field{
			Type: graphql.Int,
		},
		"end_cursor": &graphql.Field{
			Type: graphql.String,
		},
		"has_next_page": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var PageInfo_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PageInfo",
	Fields: graphql.Fields{
		"total_count": &graphql.Field{
			Type: graphql.Int,
		},
		"end_cursor": &graphql.Field{
			Type: graphql.String,
		},
		"has_next_page": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

func PageInfo_from_args(args map[string]interface{}) *PageInfo {
	return &PageInfo{}

}

var Timestamp_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Timestamp",
	Fields: graphql.Fields{
		"ISOString": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamp.Timestamp).Seconds, 0).Format("2006-01-02T15:04:05"), nil
			},
		},
		"unix": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamp.Timestamp).Seconds, 0).Unix(), nil
			},
		},
		"msSinceEpoch": &graphql.Field{
			Type:        graphql.String,
			Description: "Milliseconds since epoch (useful in JS) as a string value. Go graphql does not support int64",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				t := time.Unix(p.Source.(*timestamp.Timestamp).Seconds, 0).UnixNano()
				ms := t / int64(time.Millisecond)
				return strconv.FormatInt(ms, 10), nil
			},
		},
		"format": &graphql.Field{
			Description: `https://golang.org/pkg/time/#Time.Format Use Format() from Go's time package to format dates and times easily using the reference time "Mon Jan 2 15:04:05 -0700 MST 2006" (https://gotime.agardner.me/)`,
			Args: graphql.FieldConfigArgument{
				"layout": &graphql.ArgumentConfig{
					Description: "Mon Jan 2 15:04:05 -0700 MST 2006",
					Type:        graphql.String,
				},
			},
			Type: graphql.String,
			// Mon Jan 2 15:04:05 -0700 MST 2006
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamp.Timestamp).Seconds, 0).Format(p.Args["layout"].(string)), nil
			},
		},
	},
})

func ToTimestamp(field interface{}) *timestamp.Timestamp {
	timeMap := field.(map[string]interface{})
	t, _ := time.Parse(time.RFC3339, timeMap["ISOString"].(string))
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
	XXX_package() string
	From_args(args map[string]interface{})
}

type Svc interface {
	AppendDataloaders(map[string]Dataloader) map[string]Dataloader
}

func GrpcConnection(host string, option ...grpc.DialOption) *grpc.ClientConn {
	conn, err := grpc.Dial(
		host,
		option...,
	)

	if err != nil {
		panic(err)
	}

	return conn
}
