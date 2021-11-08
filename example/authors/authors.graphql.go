package authors

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"context"
	"github.com/graph-gophers/dataloader"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*gql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*gql.Field

var GetAuthorsRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetAuthorsRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
	},
})

var GetAuthorsRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetAuthorsRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
	},
})

var GetAuthorsRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(gql.String)),
	},
}

func GetAuthorsRequestFromArgs(args map[string]interface{}) *GetAuthorsRequest {
	return GetAuthorsRequestInstanceFromArgs(&GetAuthorsRequest{}, args)
}

func GetAuthorsRequestInstanceFromArgs(objectFromArgs *GetAuthorsRequest, args map[string]interface{}) *GetAuthorsRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		var ids []string

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	return objectFromArgs
}

func (objectFromArgs *GetAuthorsRequest) FromArgs(args map[string]interface{}) {
	GetAuthorsRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetAuthorsRequest) XXX_GraphqlType() *gql.Object {
	return GetAuthorsRequestGraphqlType
}

func (msg *GetAuthorsRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetAuthorsRequestGraphqlArgs
}

func (msg *GetAuthorsRequest) XXX_Package() string {
	return "authors"
}

var GetAuthorsResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetAuthorsResponse",
	Fields: gql.Fields{
		"authors": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(AuthorGraphqlType)),
		},
		"capitalisation1111capitalisation": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"pageInfo": &gql.Field{
			Type: pg.PageInfoGraphqlType,
		},
		"extra": &gql.Field{
			Type: extraGraphqlType,
		},
	},
})

var GetAuthorsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetAuthorsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"authors": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(AuthorGraphqlInputType)),
		},
		"capitalisation1111capitalisation": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"pageInfo": &gql.InputObjectFieldConfig{
			Type: pg.PageInfoGraphqlInputType,
		},
	},
})

var GetAuthorsResponseGraphqlArgs = gql.FieldConfigArgument{
	"authors": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(AuthorGraphqlInputType)),
	},
	"capitalisation1111capitalisation": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"pageInfo": &gql.ArgumentConfig{
		Type: pg.PageInfoGraphqlInputType,
	},
}

func GetAuthorsResponseFromArgs(args map[string]interface{}) *GetAuthorsResponse {
	return GetAuthorsResponseInstanceFromArgs(&GetAuthorsResponse{}, args)
}

func GetAuthorsResponseInstanceFromArgs(objectFromArgs *GetAuthorsResponse, args map[string]interface{}) *GetAuthorsResponse {
	if args["authors"] != nil {
		authorsInterfaceList := args["authors"].([]interface{})
		var authors []*Author

		for _, val := range authorsInterfaceList {
			itemResolved := AuthorFromArgs(val.(map[string]interface{}))
			authors = append(authors, itemResolved)
		}
		objectFromArgs.Authors = authors
	}
	if args["capitalisation1111capitalisation"] != nil {
		val := args["capitalisation1111capitalisation"]
		objectFromArgs.Capitalisation1111Capitalisation = string(val.(string))
	}
	if args["pageInfo"] != nil {
		val := args["pageInfo"]
		objectFromArgs.PageInfo = pg.PageInfoFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *GetAuthorsResponse) FromArgs(args map[string]interface{}) {
	GetAuthorsResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetAuthorsResponse) XXX_GraphqlType() *gql.Object {
	return GetAuthorsResponseGraphqlType
}

func (msg *GetAuthorsResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetAuthorsResponseGraphqlArgs
}

func (msg *GetAuthorsResponse) XXX_Package() string {
	return "authors"
}

var extraGraphqlType = gql.NewUnion(gql.UnionConfig{
	Name:  "extra",
	Types: []*gql.Object{SomeOtherThingGraphqlType, SomeThingGraphqlType},
	ResolveType: (func(p gql.ResolveTypeParams) *gql.Object {
		switch p.Value.(type) {
		case *GetAuthorsResponse_AnotherThing:
			fields := gql.Fields{}
			for name, field := range SomeOtherThingGraphqlType.Fields() {
				fields[name] = &gql.Field{
					Name: field.Name,
					Type: field.Type,
					Resolve: func(p gql.ResolveParams) (interface{}, error) {
						wrapper := p.Source.(*GetAuthorsResponse_AnotherThing)
						p.Source = wrapper.AnotherThing
						return gql.DefaultResolveFn(p)
					},
				}
			}
			return gql.NewObject(gql.ObjectConfig{
				Name:   SomeOtherThingGraphqlType.Name(),
				Fields: fields,
			})
		case *GetAuthorsResponse_Something:
			fields := gql.Fields{}
			for name, field := range SomeThingGraphqlType.Fields() {
				fields[name] = &gql.Field{
					Name: field.Name,
					Type: field.Type,
					Resolve: func(p gql.ResolveParams) (interface{}, error) {
						wrapper := p.Source.(*GetAuthorsResponse_Something)
						p.Source = wrapper.Something
						return gql.DefaultResolveFn(p)
					},
				}
			}
			return gql.NewObject(gql.ObjectConfig{
				Name:   SomeThingGraphqlType.Name(),
				Fields: fields,
			})
		}
		return nil
	}),
})

var SomeThingGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SomeThing",
	Fields: gql.Fields{
		"hello": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var SomeThingGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SomeThingInput",
	Fields: gql.InputObjectConfigFieldMap{
		"hello": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var SomeThingGraphqlArgs = gql.FieldConfigArgument{
	"hello": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func SomeThingFromArgs(args map[string]interface{}) *SomeThing {
	return SomeThingInstanceFromArgs(&SomeThing{}, args)
}

func SomeThingInstanceFromArgs(objectFromArgs *SomeThing, args map[string]interface{}) *SomeThing {
	if args["hello"] != nil {
		val := args["hello"]
		objectFromArgs.Hello = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *SomeThing) FromArgs(args map[string]interface{}) {
	SomeThingInstanceFromArgs(objectFromArgs, args)
}

func (msg *SomeThing) XXX_GraphqlType() *gql.Object {
	return SomeThingGraphqlType
}

func (msg *SomeThing) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SomeThingGraphqlArgs
}

func (msg *SomeThing) XXX_Package() string {
	return "authors"
}

var SomeOtherThingGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SomeOtherThing",
	Fields: gql.Fields{
		"world": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var SomeOtherThingGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SomeOtherThingInput",
	Fields: gql.InputObjectConfigFieldMap{
		"world": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var SomeOtherThingGraphqlArgs = gql.FieldConfigArgument{
	"world": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func SomeOtherThingFromArgs(args map[string]interface{}) *SomeOtherThing {
	return SomeOtherThingInstanceFromArgs(&SomeOtherThing{}, args)
}

func SomeOtherThingInstanceFromArgs(objectFromArgs *SomeOtherThing, args map[string]interface{}) *SomeOtherThing {
	if args["world"] != nil {
		val := args["world"]
		objectFromArgs.World = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *SomeOtherThing) FromArgs(args map[string]interface{}) {
	SomeOtherThingInstanceFromArgs(objectFromArgs, args)
}

func (msg *SomeOtherThing) XXX_GraphqlType() *gql.Object {
	return SomeOtherThingGraphqlType
}

func (msg *SomeOtherThing) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SomeOtherThingGraphqlArgs
}

func (msg *SomeOtherThing) XXX_Package() string {
	return "authors"
}

var AuthorsBatchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "AuthorsBatchRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var AuthorsBatchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AuthorsBatchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var AuthorsBatchRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func AuthorsBatchRequestFromArgs(args map[string]interface{}) *AuthorsBatchRequest {
	return AuthorsBatchRequestInstanceFromArgs(&AuthorsBatchRequest{}, args)
}

func AuthorsBatchRequestInstanceFromArgs(objectFromArgs *AuthorsBatchRequest, args map[string]interface{}) *AuthorsBatchRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		var ids []string

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	return objectFromArgs
}

func (objectFromArgs *AuthorsBatchRequest) FromArgs(args map[string]interface{}) {
	AuthorsBatchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *AuthorsBatchRequest) XXX_GraphqlType() *gql.Object {
	return AuthorsBatchRequestGraphqlType
}

func (msg *AuthorsBatchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AuthorsBatchRequestGraphqlArgs
}

func (msg *AuthorsBatchRequest) XXX_Package() string {
	return "authors"
}

var AuthorsBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "AuthorsBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var AuthorsBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AuthorsBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var AuthorsBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func AuthorsBatchResponseFromArgs(args map[string]interface{}) *AuthorsBatchResponse {
	return AuthorsBatchResponseInstanceFromArgs(&AuthorsBatchResponse{}, args)
}

func AuthorsBatchResponseInstanceFromArgs(objectFromArgs *AuthorsBatchResponse, args map[string]interface{}) *AuthorsBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *AuthorsBatchResponse) FromArgs(args map[string]interface{}) {
	AuthorsBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *AuthorsBatchResponse) XXX_GraphqlType() *gql.Object {
	return AuthorsBatchResponseGraphqlType
}

func (msg *AuthorsBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AuthorsBatchResponseGraphqlArgs
}

func (msg *AuthorsBatchResponse) XXX_Package() string {
	return "authors"
}

var AuthorsBoolBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "AuthorsBoolBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var AuthorsBoolBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AuthorsBoolBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var AuthorsBoolBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func AuthorsBoolBatchResponseFromArgs(args map[string]interface{}) *AuthorsBoolBatchResponse {
	return AuthorsBoolBatchResponseInstanceFromArgs(&AuthorsBoolBatchResponse{}, args)
}

func AuthorsBoolBatchResponseInstanceFromArgs(objectFromArgs *AuthorsBoolBatchResponse, args map[string]interface{}) *AuthorsBoolBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *AuthorsBoolBatchResponse) FromArgs(args map[string]interface{}) {
	AuthorsBoolBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *AuthorsBoolBatchResponse) XXX_GraphqlType() *gql.Object {
	return AuthorsBoolBatchResponseGraphqlType
}

func (msg *AuthorsBoolBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AuthorsBoolBatchResponseGraphqlArgs
}

func (msg *AuthorsBoolBatchResponse) XXX_Package() string {
	return "authors"
}

var AuthorGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Author",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var AuthorGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AuthorInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var AuthorGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func AuthorFromArgs(args map[string]interface{}) *Author {
	return AuthorInstanceFromArgs(&Author{}, args)
}

func AuthorInstanceFromArgs(objectFromArgs *Author, args map[string]interface{}) *Author {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["name"] != nil {
		val := args["name"]
		objectFromArgs.Name = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *Author) FromArgs(args map[string]interface{}) {
	AuthorInstanceFromArgs(objectFromArgs, args)
}

func (msg *Author) XXX_GraphqlType() *gql.Object {
	return AuthorGraphqlType
}

func (msg *Author) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AuthorGraphqlArgs
}

func (msg *Author) XXX_Package() string {
	return "authors"
}

var AuthorsClientInstance AuthorsClient

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		AuthorsClientInstance = NewAuthorsClient(pg.GrpcConnection("localhost:50052", opts...))
	})
	fields = append(fields, &gql.Field{
		Name: "authors_GetAuthors",
		Type: GetAuthorsResponseGraphqlType,
		Args: GetAuthorsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return AuthorsClientInstance.GetAuthors(p.Context, GetAuthorsRequestFromArgs(p.Args))
		},
	})

}

func WithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "LoadAuthorsLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			resp, err := AuthorsClientInstance.LoadAuthors(ctx, &pg.BatchRequest{
				Keys: keys.Keys(),
			})

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *Author
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	ctx = context.WithValue(ctx, "LoadAuthorsBoolLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			resp, err := AuthorsClientInstance.LoadAuthorsBool(ctx, &pg.BatchRequest{
				Keys: keys.Keys(),
			})

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty bool
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	return ctx
}

func LoadAuthors(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadAuthorsLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadAuthorsLoader").(*dataloader.Loader)
	default:
		panic("Please call authors.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*Author), nil
	}, nil
}

func LoadAuthorsMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadAuthorsLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadAuthorsLoader").(*dataloader.Loader)
	default:
		panic("Please call authors.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*Author
		for _, res := range resSlice {
			results = append(results, res.(*Author))
		}

		return results, nil
	}, nil
}

func LoadAuthorsBool(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadAuthorsBoolLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadAuthorsBoolLoader").(*dataloader.Loader)
	default:
		panic("Please call authors.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(bool), nil
	}, nil
}

func LoadAuthorsBoolMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadAuthorsBoolLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadAuthorsBoolLoader").(*dataloader.Loader)
	default:
		panic("Please call authors.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []bool
		for _, res := range resSlice {
			results = append(results, res.(bool))
		}

		return results, nil
	}, nil
}
