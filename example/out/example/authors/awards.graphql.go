package authors

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

var AwardGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Award",
	Fields: gql.Fields{
		"title": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"year": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var AwardGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AwardInput",
	Fields: gql.InputObjectConfigFieldMap{
		"title": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"year": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var AwardGraphqlArgs = gql.FieldConfigArgument{
	"title": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"year": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func AwardFromArgs(args map[string]interface{}) *Award {
	return AwardInstanceFromArgs(&Award{}, args)
}

func AwardInstanceFromArgs(objectFromArgs *Award, args map[string]interface{}) *Award {
	if args["title"] != nil {
		val := args["title"]
		objectFromArgs.Title = string(val.(string))
	}
	if args["year"] != nil {
		val := args["year"]
		objectFromArgs.Year = int64(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *Award) FromArgs(args map[string]interface{}) {
	AwardInstanceFromArgs(objectFromArgs, args)
}

func (msg *Award) XXX_GraphqlType() *gql.Object {
	return AwardGraphqlType
}

func (msg *Award) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AwardGraphqlArgs
}

func (msg *Award) XXX_Package() string {
	return "authors"
}
