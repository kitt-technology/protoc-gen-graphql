package cases

import (
	"github.com/graphql-go/graphql"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var Fields []*graphql.Field

var GetSomethingRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetSomethingRequest",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var GetSomethingRequest_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
}

func GetSomethingRequest_from_args(args map[string]interface{}) *GetSomethingRequest {
	objectFromArgs := GetSomethingRequest{}
	if args["id"] != nil {

		idInterfaceList := args["id"].([]interface{})

		var id []string
		for _, item := range idInterfaceList {
			id = append(id, item.(string))
		}
		objectFromArgs.Id = id

	}

	return &objectFromArgs
}

var GetSomethingResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetSomethingResponse",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"somethingId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"someIdList": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var GetSomethingResponse_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"somethingId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"someIdList": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
}

func GetSomethingResponse_from_args(args map[string]interface{}) *GetSomethingResponse {
	objectFromArgs := GetSomethingResponse{}

	objectFromArgs.Id = args["id"].(string)

	objectFromArgs.SomethingId = args["somethingId"].(string)

	if args["someIdList"] != nil {

		someIdListInterfaceList := args["someIdList"].([]interface{})

		var someIdList []string
		for _, item := range someIdListInterfaceList {
			someIdList = append(someIdList, item.(string))
		}
		objectFromArgs.SomeIdList = someIdList

	}

	return &objectFromArgs
}

var TestCommand_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "TestCommand",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"someInt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"someBool": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"idObjectList": &graphql.Field{
			Type: graphql.NewList(SomeSmallObject_type),
		},
		"stringList": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"optionalId": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var TestCommand_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"someInt": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"someBool": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Boolean),
	},
	"idObjectList": &graphql.ArgumentConfig{
		Type: graphql.NewList(SomeSmallObject_type),
	},
	"stringList": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
	"optionalId": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func TestCommand_from_args(args map[string]interface{}) *TestCommand {
	objectFromArgs := TestCommand{}

	objectFromArgs.Id = args["id"].(string)

	objectFromArgs.SomeInt = args["someInt"].(int32)

	objectFromArgs.SomeBool = args["someBool"].(bool)

	if args["idObjectList"] != nil {

		idObjectListInterfaceList := args["idObjectList"].([]interface{})

		var idObjectList []*SomeSmallObject
		for _, item := range idObjectListInterfaceList {
			idObjectList = append(idObjectList, item.(*SomeSmallObject))
		}
		objectFromArgs.IdObjectList = idObjectList

	}
	if args["stringList"] != nil {

		stringListInterfaceList := args["stringList"].([]interface{})

		var stringList []string
		for _, item := range stringListInterfaceList {
			stringList = append(stringList, item.(string))
		}
		objectFromArgs.StringList = stringList

	}

	if args["optionalId"] != nil {
		objectFromArgs.OptionalId = wrapperspb.String(args["optionalId"].(string))
	}

	return &objectFromArgs
}

var SomeSmallObject_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "SomeSmallObject",
	Fields: graphql.Fields{
		"someId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var SomeSmallObject_args = graphql.FieldConfigArgument{
	"someId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func SomeSmallObject_from_args(args map[string]interface{}) *SomeSmallObject {
	objectFromArgs := SomeSmallObject{}

	objectFromArgs.SomeId = args["someId"].(string)

	return &objectFromArgs
}

var client MyServiceClient

func init() {
	client = NewMyServiceClient(pg.GrpcConnection(""))
	Fields = append(Fields, &graphql.Field{
		Name: "MyService_GetSomething",
		Type: GetSomethingResponse_type,
		Args: GetSomethingRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return client.GetSomething(p.Context, GetSomethingRequest_from_args(p.Args))
		},
	})

}
