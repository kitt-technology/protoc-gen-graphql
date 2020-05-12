package cases

import "github.com/graphql-go/graphql"

var TestCommand_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "TestCommand",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
		"someInt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.Int)),
		},
		"someBool": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.Boolean)),
		},
		"idObjectList": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewList(SomeObject_type))),
		},
		"stringList": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewList(graphql.String))),
		},
		"optionalId": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SomeObject_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "SomeObject",
	Fields: graphql.Fields{
		"someId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
	},
})

var CommandSuccess_type = graphql.NewObject(graphql.ObjectConfig{
	Name:   "CommandSuccess",
	Fields: graphql.Fields{},
})

var CommandFailed_type = graphql.NewObject(graphql.ObjectConfig{
	Name:   "CommandFailed",
	Fields: graphql.Fields{},
})
