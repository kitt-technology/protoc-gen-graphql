package example

import (
	gql "github.com/graphql-go/graphql"
	"context"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
)

var SearchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SearchRequest",
	Fields: gql.Fields{
		"query": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"limit": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"offset": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})
var SearchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SearchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"query": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"limit": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"offset": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var SearchRequestGraphqlArgs = gql.FieldConfigArgument{
	"query": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"limit": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"offset": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func SearchRequestFromArgs(args map[string]interface{}) *SearchRequest {
	return SearchRequestInstanceFromArgs(&SearchRequest{}, args)
}

func SearchRequestInstanceFromArgs(objectFromArgs *SearchRequest, args map[string]interface{}) *SearchRequest {
	if args["query"] != nil {
		val := args["query"]
		objectFromArgs.Query = string(val.(string))
	}
	if args["limit"] != nil {
		val := args["limit"]
		objectFromArgs.Limit = int32(val.(int))
	}
	if args["offset"] != nil {
		val := args["offset"]
		objectFromArgs.Offset = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *SearchRequest) FromArgs(args map[string]interface{}) {
	SearchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *SearchRequest) XXX_GraphqlType() *gql.Object {
	return SearchRequestGraphqlType
}

func (msg *SearchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SearchRequestGraphqlArgs
}

func (msg *SearchRequest) XXX_Package() string {
	return "example"
}

var SearchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SearchResponse",
	Fields: gql.Fields{
		"results": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"total": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})
var SearchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SearchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"results": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"total": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var SearchResponseGraphqlArgs = gql.FieldConfigArgument{
	"results": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
	"total": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func SearchResponseFromArgs(args map[string]interface{}) *SearchResponse {
	return SearchResponseInstanceFromArgs(&SearchResponse{}, args)
}

func SearchResponseInstanceFromArgs(objectFromArgs *SearchResponse, args map[string]interface{}) *SearchResponse {
	if args["results"] != nil {
		resultsInterfaceList := args["results"].([]interface{})
		results := make([]string, 0)

		for _, val := range resultsInterfaceList {
			itemResolved := string(val.(string))
			results = append(results, itemResolved)
		}
		objectFromArgs.Results = results
	}
	if args["total"] != nil {
		val := args["total"]
		objectFromArgs.Total = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *SearchResponse) FromArgs(args map[string]interface{}) {
	SearchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *SearchResponse) XXX_GraphqlType() *gql.Object {
	return SearchResponseGraphqlType
}

func (msg *SearchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SearchResponseGraphqlArgs
}

func (msg *SearchResponse) XXX_Package() string {
	return "example"
}

// allMessages contains all message types from this proto package
var allMessages = []pg.GraphqlMessage{
	&SearchRequest{},
	&SearchResponse{},
}

// ExampleModule implements the Module interface for the example package (types only, no services)
type ExampleModule struct{}

// NewExampleModule creates a new module instance
func NewExampleModule() pg.Module {
	return &ExampleModule{}
}

// Fields returns an empty map (no services in this module)
func (m *ExampleModule) Fields() gql.Fields {
	return gql.Fields{}
}

// Messages returns all message types from this package
func (m *ExampleModule) Messages() []pg.GraphqlMessage {
	return allMessages
}

// WithLoaders returns the context unchanged (no loaders in this module)
func (m *ExampleModule) WithLoaders(ctx context.Context) context.Context {
	return ctx
}

// PackageName returns the proto package name
func (m *ExampleModule) PackageName() string {
	return "example"
}
