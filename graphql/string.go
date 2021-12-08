package graphql

import (
	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var WrappedString = gql.NewScalar(gql.ScalarConfig{
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
