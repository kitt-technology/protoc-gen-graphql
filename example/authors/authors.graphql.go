package authors

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"context"
	"github.com/graph-gophers/dataloader"
	"os"
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
		ids := make([]string, 0)

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
		"pageInfo": &gql.Field{
			Type: pg.PageInfoGraphqlType,
		},
	},
})

var GetAuthorsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetAuthorsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"authors": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(AuthorGraphqlInputType)),
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
		authors := make([]*Author, 0)

		for _, val := range authorsInterfaceList {
			itemResolved := AuthorFromArgs(val.(map[string]interface{}))
			authors = append(authors, itemResolved)
		}
		objectFromArgs.Authors = authors
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
var AuthorsServiceInstance AuthorsServer
var AuthorsDialOpts []grpc.DialOption

func init() {
	fields = append(fields, &gql.Field{
		Name: "authors_GetAuthors",
		Type: GetAuthorsResponseGraphqlType,
		Args: GetAuthorsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if AuthorsServiceInstance != nil {
				return AuthorsServiceInstance.GetAuthors(p.Context, GetAuthorsRequestFromArgs(p.Args))
			}
			if AuthorsClientInstance == nil {
				AuthorsClientInstance = getAuthorsClient()
			}
			return AuthorsClientInstance.GetAuthors(p.Context, GetAuthorsRequestFromArgs(p.Args))
		},
	})

}

func getAuthorsClient() AuthorsClient {
	host := "localhost:50052"
	envHost := os.Getenv("SERVICE_HOST")
	if envHost != "" {
		host = envHost
	}
	return NewAuthorsClient(pg.GrpcConnection(host, AuthorsDialOpts...))
}

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		AuthorsDialOpts = opts
	})
}

// SetAuthorsService sets the service implementation for direct calls (no gRPC)
func SetAuthorsService(service AuthorsServer) {
	AuthorsServiceInstance = service
}

// SetAuthorsClient sets the gRPC client for remote calls
func SetAuthorsClient(client AuthorsClient) {
	AuthorsClientInstance = client
}

func WithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "LoadAuthorsLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *pg.BatchResponse
			var err error
			if AuthorsServiceInstance != nil {
				resp, err = AuthorsServiceInstance.LoadAuthors(ctx, &pg.BatchRequest{
					Keys: keys.Keys(),
				})
			} else {
				if AuthorsClientInstance == nil {
					AuthorsClientInstance = getAuthorsClient()
				}
				resp, err = AuthorsClientInstance.LoadAuthors(ctx, &pg.BatchRequest{
					Keys: keys.Keys(),
				})
			}

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
