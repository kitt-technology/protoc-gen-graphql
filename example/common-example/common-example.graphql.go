package common_example

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

var Int32RangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Int32Range",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"max": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var Int32RangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "Int32RangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"max": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var Int32RangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"max": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func Int32RangeFromArgs(args map[string]interface{}) *Int32Range {
	return Int32RangeInstanceFromArgs(&Int32Range{}, args)
}

func Int32RangeInstanceFromArgs(objectFromArgs *Int32Range, args map[string]interface{}) *Int32Range {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = int32(val.(int))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *Int32Range) FromArgs(args map[string]interface{}) {
	Int32RangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *Int32Range) XXX_GraphqlType() *gql.Object {
	return Int32RangeGraphqlType
}

func (msg *Int32Range) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return Int32RangeGraphqlArgs
}

func (msg *Int32Range) XXX_Package() string {
	return "common_example"
}

var MoneyRangeGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "MoneyRange",
	Fields: gql.Fields{
		"min": &gql.Field{
			Type: MoneyGraphqlType,
		},
		"max": &gql.Field{
			Type: MoneyGraphqlType,
		},
	},
})

var MoneyRangeGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "MoneyRangeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"min": &gql.InputObjectFieldConfig{
			Type: MoneyGraphqlInputType,
		},
		"max": &gql.InputObjectFieldConfig{
			Type: MoneyGraphqlInputType,
		},
	},
})

var MoneyRangeGraphqlArgs = gql.FieldConfigArgument{
	"min": &gql.ArgumentConfig{
		Type: MoneyGraphqlInputType,
	},
	"max": &gql.ArgumentConfig{
		Type: MoneyGraphqlInputType,
	},
}

func MoneyRangeFromArgs(args map[string]interface{}) *MoneyRange {
	return MoneyRangeInstanceFromArgs(&MoneyRange{}, args)
}

func MoneyRangeInstanceFromArgs(objectFromArgs *MoneyRange, args map[string]interface{}) *MoneyRange {
	if args["min"] != nil {
		val := args["min"]
		objectFromArgs.Min = MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["max"] != nil {
		val := args["max"]
		objectFromArgs.Max = MoneyFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *MoneyRange) FromArgs(args map[string]interface{}) {
	MoneyRangeInstanceFromArgs(objectFromArgs, args)
}

func (msg *MoneyRange) XXX_GraphqlType() *gql.Object {
	return MoneyRangeGraphqlType
}

func (msg *MoneyRange) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return MoneyRangeGraphqlArgs
}

func (msg *MoneyRange) XXX_Package() string {
	return "common_example"
}
