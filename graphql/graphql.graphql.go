package graphql

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

var BatchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "BatchRequest",
	Fields: gql.Fields{
		"keys": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var BatchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BatchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"keys": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var BatchRequestGraphqlArgs = gql.FieldConfigArgument{
	"keys": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func BatchRequestFromArgs(args map[string]interface{}) *BatchRequest {
	return BatchRequestInstanceFromArgs(&BatchRequest{}, args)
}

func BatchRequestInstanceFromArgs(objectFromArgs *BatchRequest, args map[string]interface{}) *BatchRequest {
	if args["keys"] != nil {
		keysInterfaceList := args["keys"].([]interface{})
		var keys []string

		for _, val := range keysInterfaceList {
			itemResolved := string(val.(string))
			keys = append(keys, itemResolved)
		}
		objectFromArgs.Keys = keys
	}
	return objectFromArgs
}

func (objectFromArgs *BatchRequest) FromArgs(args map[string]interface{}) {
	BatchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *BatchRequest) XXX_GraphqlType() *gql.Object {
	return BatchRequestGraphqlType
}

func (msg *BatchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return BatchRequestGraphqlArgs
}

func (msg *BatchRequest) XXX_Package() string {
	return "graphql"
}

var PageInfoGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "PageInfo",
	Fields: gql.Fields{
		"totalCount": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"endCursor": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"hasNextPage": &gql.Field{
			Type: gql.NewNonNull(gql.Boolean),
		},
	},
})

var PageInfoGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "PageInfoInput",
	Fields: gql.InputObjectConfigFieldMap{
		"totalCount": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"endCursor": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"hasNextPage": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Boolean),
		},
	},
})

var PageInfoGraphqlArgs = gql.FieldConfigArgument{
	"totalCount": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"endCursor": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"hasNextPage": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Boolean),
	},
}

func PageInfoFromArgs(args map[string]interface{}) *PageInfo {
	return PageInfoInstanceFromArgs(&PageInfo{}, args)
}

func PageInfoInstanceFromArgs(objectFromArgs *PageInfo, args map[string]interface{}) *PageInfo {
	if args["totalCount"] != nil {
		val := args["totalCount"]
		objectFromArgs.TotalCount = int32(val.(int))
	}
	if args["endCursor"] != nil {
		val := args["endCursor"]
		objectFromArgs.EndCursor = string(val.(string))
	}
	if args["hasNextPage"] != nil {
		val := args["hasNextPage"]
		objectFromArgs.HasNextPage = bool(val.(bool))
	}
	return objectFromArgs
}

func (objectFromArgs *PageInfo) FromArgs(args map[string]interface{}) {
	PageInfoInstanceFromArgs(objectFromArgs, args)
}

func (msg *PageInfo) XXX_GraphqlType() *gql.Object {
	return PageInfoGraphqlType
}

func (msg *PageInfo) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return PageInfoGraphqlArgs
}

func (msg *PageInfo) XXX_Package() string {
	return "graphql"
}

var FieldMaskGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "FieldMask",
	Fields: gql.Fields{
		"paths": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var FieldMaskGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "FieldMaskInput",
	Fields: gql.InputObjectConfigFieldMap{
		"paths": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var FieldMaskGraphqlArgs = gql.FieldConfigArgument{
	"paths": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func FieldMaskFromArgs(args map[string]interface{}) *FieldMask {
	return FieldMaskInstanceFromArgs(&FieldMask{}, args)
}

func FieldMaskInstanceFromArgs(objectFromArgs *FieldMask, args map[string]interface{}) *FieldMask {
	if args["paths"] != nil {
		pathsInterfaceList := args["paths"].([]interface{})
		var paths []string

		for _, val := range pathsInterfaceList {
			itemResolved := string(val.(string))
			paths = append(paths, itemResolved)
		}
		objectFromArgs.Paths = paths
	}
	return objectFromArgs
}

func (objectFromArgs *FieldMask) FromArgs(args map[string]interface{}) {
	FieldMaskInstanceFromArgs(objectFromArgs, args)
}

func (msg *FieldMask) XXX_GraphqlType() *gql.Object {
	return FieldMaskGraphqlType
}

func (msg *FieldMask) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return FieldMaskGraphqlArgs
}

func (msg *FieldMask) XXX_Package() string {
	return "graphql"
}
