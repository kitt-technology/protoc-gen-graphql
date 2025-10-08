package graphql

import (
	"strconv"
	"time"

	gql "github.com/graphql-go/graphql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTimestamp(field interface{}) *timestamppb.Timestamp {
	timeMap, ok := field.(map[string]interface{})
	if !ok {
		return nil
	}
	isoString, ok := timeMap["ISOString"].(string)
	if !ok {
		return nil
	}
	t, err := time.Parse(time.RFC3339, isoString)
	if err != nil {
		return nil
	}
	ts := timestamppb.Timestamp{
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
				if ts, ok := p.Source.(*timestamppb.Timestamp); ok {
					return time.Unix(ts.Seconds, 0).Format(time.RFC3339), nil
				}
				return nil, nil
			},
		},
		"unix": &gql.Field{
			Type: gql.Int,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if ts, ok := p.Source.(*timestamppb.Timestamp); ok {
					return time.Unix(ts.Seconds, 0).Unix(), nil
				}
				return nil, nil
			},
		},
		"msSinceEpoch": &gql.Field{
			Type:        gql.String,
			Description: "Milliseconds since epoch (useful in JS) as a string value. Go graphql does not support int64",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if ts, ok := p.Source.(*timestamppb.Timestamp); ok {
					t := time.Unix(ts.Seconds, 0).UnixNano()
					ms := t / int64(time.Millisecond)
					return strconv.FormatInt(ms, 10), nil
				}
				return nil, nil
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
				if ts, ok := p.Source.(*timestamppb.Timestamp); ok {
					if layout, ok := p.Args["layout"].(string); ok {
						return time.Unix(ts.Seconds, 0).Format(layout), nil
					}
				}
				return nil, nil
			},
		},
	},
})
