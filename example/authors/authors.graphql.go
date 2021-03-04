package authors

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"

	"context"

	"github.com/graph-gophers/dataloader"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*graphql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*graphql.Field

var GetAuthorsRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
		},
	},
})

var GetAuthorsRequest_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetAuthorsRequestInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
		},
	},
})

var GetAuthorsRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
	},
}

func GetAuthorsRequest_from_args(args map[string]interface{}) *GetAuthorsRequest {
	return GetAuthorsRequest_instance_from_args(&GetAuthorsRequest{}, args)
}

func GetAuthorsRequest_instance_from_args(objectFromArgs *GetAuthorsRequest, args map[string]interface{}) *GetAuthorsRequest {
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

func (objectFromArgs *GetAuthorsRequest) From_args(args map[string]interface{}) {
	GetAuthorsRequest_instance_from_args(objectFromArgs, args)
}

func (msg *GetAuthorsRequest) XXX_type() *graphql.Object {
	return GetAuthorsRequest_type
}

func (msg *GetAuthorsRequest) XXX_args() graphql.FieldConfigArgument {
	return GetAuthorsRequest_args
}

func (msg *GetAuthorsRequest) XXX_package() string {
	return "authors"
}

var GetAuthorsResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsResponse",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(Author_type)),
		},
		"capitalisation1111capitalisation": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"extra": &graphql.Field{
			Type: extra_type,
		},
	},
})

var GetAuthorsResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetAuthorsResponseInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"authors": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(Author_input_type)),
		},
		"capitalisation1111capitalisation": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var GetAuthorsResponse_args = graphql.FieldConfigArgument{
	"authors": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(Author_input_type)),
	},
	"capitalisation1111capitalisation": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func GetAuthorsResponse_from_args(args map[string]interface{}) *GetAuthorsResponse {
	return GetAuthorsResponse_instance_from_args(&GetAuthorsResponse{}, args)
}

func GetAuthorsResponse_instance_from_args(objectFromArgs *GetAuthorsResponse, args map[string]interface{}) *GetAuthorsResponse {
	if args["authors"] != nil {

		authorsInterfaceList := args["authors"].([]interface{})

		var authors []*Author

		for _, val := range authorsInterfaceList {
			itemResolved := Author_from_args(val.(map[string]interface{}))
			authors = append(authors, itemResolved)
		}
		objectFromArgs.Authors = authors
	}
	if args["capitalisation1111capitalisation"] != nil {
		val := args["capitalisation1111capitalisation"]
		objectFromArgs.Capitalisation1111Capitalisation = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *GetAuthorsResponse) From_args(args map[string]interface{}) {
	GetAuthorsResponse_instance_from_args(objectFromArgs, args)
}

func (msg *GetAuthorsResponse) XXX_type() *graphql.Object {
	return GetAuthorsResponse_type
}

func (msg *GetAuthorsResponse) XXX_args() graphql.FieldConfigArgument {
	return GetAuthorsResponse_args
}

func (msg *GetAuthorsResponse) XXX_package() string {
	return "authors"
}

var extra_type = graphql.NewUnion(graphql.UnionConfig{
	Name:  "extra",
	Types: []*graphql.Object{SomeOtherThing_type, SomeThing_type},
	ResolveType: (func(p graphql.ResolveTypeParams) *graphql.Object {
		switch p.Value.(type) {
		case *GetAuthorsResponse_AnotherThing:
			fields := graphql.Fields{}
			for name, field := range SomeOtherThing_type.Fields() {
				fields[name] = &graphql.Field{
					Name: field.Name,
					Type: field.Type,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						wrapper := p.Source.(*GetAuthorsResponse_AnotherThing)
						p.Source = wrapper.AnotherThing
						return graphql.DefaultResolveFn(p)
					},
				}
			}
			return graphql.NewObject(graphql.ObjectConfig{
				Name:   SomeOtherThing_type.Name(),
				Fields: fields,
			})
		case *GetAuthorsResponse_Something:
			fields := graphql.Fields{}
			for name, field := range SomeThing_type.Fields() {
				fields[name] = &graphql.Field{
					Name: field.Name,
					Type: field.Type,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						wrapper := p.Source.(*GetAuthorsResponse_Something)
						p.Source = wrapper.Something
						return graphql.DefaultResolveFn(p)
					},
				}
			}
			return graphql.NewObject(graphql.ObjectConfig{
				Name:   SomeThing_type.Name(),
				Fields: fields,
			})
		}
		return nil
	}),
})

var SomeThing_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "SomeThing",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var SomeThing_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SomeThingInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"hello": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var SomeThing_args = graphql.FieldConfigArgument{
	"hello": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func SomeThing_from_args(args map[string]interface{}) *SomeThing {
	return SomeThing_instance_from_args(&SomeThing{}, args)
}

func SomeThing_instance_from_args(objectFromArgs *SomeThing, args map[string]interface{}) *SomeThing {
	if args["hello"] != nil {
		val := args["hello"]
		objectFromArgs.Hello = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *SomeThing) From_args(args map[string]interface{}) {
	SomeThing_instance_from_args(objectFromArgs, args)
}

func (msg *SomeThing) XXX_type() *graphql.Object {
	return SomeThing_type
}

func (msg *SomeThing) XXX_args() graphql.FieldConfigArgument {
	return SomeThing_args
}

func (msg *SomeThing) XXX_package() string {
	return "authors"
}

var SomeOtherThing_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "SomeOtherThing",
	Fields: graphql.Fields{
		"world": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var SomeOtherThing_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "SomeOtherThingInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"world": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var SomeOtherThing_args = graphql.FieldConfigArgument{
	"world": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func SomeOtherThing_from_args(args map[string]interface{}) *SomeOtherThing {
	return SomeOtherThing_instance_from_args(&SomeOtherThing{}, args)
}

func SomeOtherThing_instance_from_args(objectFromArgs *SomeOtherThing, args map[string]interface{}) *SomeOtherThing {
	if args["world"] != nil {
		val := args["world"]
		objectFromArgs.World = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *SomeOtherThing) From_args(args map[string]interface{}) {
	SomeOtherThing_instance_from_args(objectFromArgs, args)
}

func (msg *SomeOtherThing) XXX_type() *graphql.Object {
	return SomeOtherThing_type
}

func (msg *SomeOtherThing) XXX_args() graphql.FieldConfigArgument {
	return SomeOtherThing_args
}

func (msg *SomeOtherThing) XXX_package() string {
	return "authors"
}

var AuthorsBatchRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthorsBatchRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},
})

var AuthorsBatchRequest_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorsBatchRequestInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
		},
	},
})

var AuthorsBatchRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),
	},
}

func AuthorsBatchRequest_from_args(args map[string]interface{}) *AuthorsBatchRequest {
	return AuthorsBatchRequest_instance_from_args(&AuthorsBatchRequest{}, args)
}

func AuthorsBatchRequest_instance_from_args(objectFromArgs *AuthorsBatchRequest, args map[string]interface{}) *AuthorsBatchRequest {
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

func (objectFromArgs *AuthorsBatchRequest) From_args(args map[string]interface{}) {
	AuthorsBatchRequest_instance_from_args(objectFromArgs, args)
}

func (msg *AuthorsBatchRequest) XXX_type() *graphql.Object {
	return AuthorsBatchRequest_type
}

func (msg *AuthorsBatchRequest) XXX_args() graphql.FieldConfigArgument {
	return AuthorsBatchRequest_args
}

func (msg *AuthorsBatchRequest) XXX_package() string {
	return "authors"
}

var AuthorsBatchResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthorsBatchResponse",
	Fields: graphql.Fields{
		"_null": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var AuthorsBatchResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorsBatchResponseInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"_null": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})

var AuthorsBatchResponse_args = graphql.FieldConfigArgument{
	"_null": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
}

func AuthorsBatchResponse_from_args(args map[string]interface{}) *AuthorsBatchResponse {
	return AuthorsBatchResponse_instance_from_args(&AuthorsBatchResponse{}, args)
}

func AuthorsBatchResponse_instance_from_args(objectFromArgs *AuthorsBatchResponse, args map[string]interface{}) *AuthorsBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *AuthorsBatchResponse) From_args(args map[string]interface{}) {
	AuthorsBatchResponse_instance_from_args(objectFromArgs, args)
}

func (msg *AuthorsBatchResponse) XXX_type() *graphql.Object {
	return AuthorsBatchResponse_type
}

func (msg *AuthorsBatchResponse) XXX_args() graphql.FieldConfigArgument {
	return AuthorsBatchResponse_args
}

func (msg *AuthorsBatchResponse) XXX_package() string {
	return "authors"
}

var Author_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var Author_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AuthorInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var Author_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func Author_from_args(args map[string]interface{}) *Author {
	return Author_instance_from_args(&Author{}, args)
}

func Author_instance_from_args(objectFromArgs *Author, args map[string]interface{}) *Author {
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

func (objectFromArgs *Author) From_args(args map[string]interface{}) {
	Author_instance_from_args(objectFromArgs, args)
}

func (msg *Author) XXX_type() *graphql.Object {
	return Author_type
}

func (msg *Author) XXX_args() graphql.FieldConfigArgument {
	return Author_args
}

func (msg *Author) XXX_package() string {
	return "authors"
}

var AuthorsClientInstance AuthorsClient

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		AuthorsClientInstance = NewAuthorsClient(pg.GrpcConnection("localhost:50052", opts...))
	})
	fields = append(fields, &graphql.Field{
		Name: "authors_GetAuthors",
		Type: GetAuthorsResponse_type,
		Args: GetAuthorsRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return AuthorsClientInstance.GetAuthors(p.Context, GetAuthorsRequest_from_args(p.Args))
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
				results = append(results, &dataloader.Result{Data: resp.Results[key]})
			}

			return results
		},
	))

	return ctx
}

func LoadAuthors(p graphql.ResolveParams, key string) (func() (interface{}, error), error) {
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

func LoadAuthorsMany(p graphql.ResolveParams, keys []string) (func() (interface{}, error), error) {
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
