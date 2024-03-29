package graphql

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	gql "github.com/graphql-go/graphql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func ToTimestamp(field interface{}) *timestamp.Timestamp {
	timeMap := field.(map[string]interface{})
	t, _ := time.Parse(time.RFC3339, timeMap["ISOString"].(string))
	ts := timestamp.Timestamp{
		Seconds: t.Unix(),
	}
	return &ts
}

var TimestampGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "TimestampInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ISOString": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
	},
})

var TimestampGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Timestamp",
	Fields: gql.Fields{
		"ISOString": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamppb.Timestamp).Seconds, 0).Format(time.RFC3339), nil
			},
		},
		"unix": &gql.Field{
			Type: gql.Int,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamppb.Timestamp).Seconds, 0).Unix(), nil
			},
		},
		"msSinceEpoch": &gql.Field{
			Type:        gql.String,
			Description: "Milliseconds since epoch (useful in JS) as a string value. Go graphql does not support int64",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				t := time.Unix(p.Source.(*timestamppb.Timestamp).Seconds, 0).UnixNano()
				ms := t / int64(time.Millisecond)
				return strconv.FormatInt(ms, 10), nil
			},
		},
		"format": &gql.Field{
			Description: `https://golang.org/pkg/time/#Time.Format Use Format() from Go's time package to format dates and times easily using the reference time "Mon Jan 2 15:04:05 -0700 MST 2006" (https://gotime.agardner.me/)`,
			Args: gql.FieldConfigArgument{
				"layout": &gql.ArgumentConfig{
					Description: "Mon Jan 2 15:04:05 -0700 MST 2006",
					Type:        gql.String,
				},
			},
			Type: gql.String,
			// Mon Jan 2 15:04:05 -0700 MST 2006
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return time.Unix(p.Source.(*timestamppb.Timestamp).Seconds, 0).Format(p.Args["layout"].(string)), nil
			},
		},
	},
})
