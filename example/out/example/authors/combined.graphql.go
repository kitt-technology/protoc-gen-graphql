package authors

import (
	gql "github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql/language/ast"
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

var AwardTypeGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "AwardType",
	Values: gql.EnumValueConfigMap{
		"BestModernFiction": &gql.EnumValueConfig{
			Value: AwardType(2),
		},
		"BookOfTheYear": &gql.EnumValueConfig{
			Value: AwardType(1),
		},
		"UnknownAwardType": &gql.EnumValueConfig{
			Value: AwardType(0),
		},
	},
})

var AwardTypeGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "AwardType",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(AwardType).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var AwardGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Award",
	Fields: gql.Fields{
		"title": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"year": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"importance": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"type": &gql.Field{
			Type: AwardTypeGraphqlEnum,
		},
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
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
		"importance": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"type": &gql.InputObjectFieldConfig{
			Type: AwardTypeGraphqlEnum,
		},
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
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
	"importance": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"type": &gql.ArgumentConfig{
		Type: AwardTypeGraphqlEnum,
	},
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
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
	if args["importance"] != nil {
		val := args["importance"]
		objectFromArgs.Importance = int64(val.(int))
	}
	if args["type"] != nil {
		val := args["type"]
		objectFromArgs.Type = val.(AwardType)
	}
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
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

var GetAwardsRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetAwardsRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var GetAwardsRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetAwardsRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var GetAwardsRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func GetAwardsRequestFromArgs(args map[string]interface{}) *GetAwardsRequest {
	return GetAwardsRequestInstanceFromArgs(&GetAwardsRequest{}, args)
}

func GetAwardsRequestInstanceFromArgs(objectFromArgs *GetAwardsRequest, args map[string]interface{}) *GetAwardsRequest {
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

func (objectFromArgs *GetAwardsRequest) FromArgs(args map[string]interface{}) {
	GetAwardsRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetAwardsRequest) XXX_GraphqlType() *gql.Object {
	return GetAwardsRequestGraphqlType
}

func (msg *GetAwardsRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetAwardsRequestGraphqlArgs
}

func (msg *GetAwardsRequest) XXX_Package() string {
	return "authors"
}

var GetAwardsResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetAwardsResponse",
	Fields: gql.Fields{
		"awards": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(AwardGraphqlType)),
		},
	},
})

var GetAwardsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetAwardsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"awards": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(AwardGraphqlInputType)),
		},
	},
})

var GetAwardsResponseGraphqlArgs = gql.FieldConfigArgument{
	"awards": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(AwardGraphqlInputType)),
	},
}

func GetAwardsResponseFromArgs(args map[string]interface{}) *GetAwardsResponse {
	return GetAwardsResponseInstanceFromArgs(&GetAwardsResponse{}, args)
}

func GetAwardsResponseInstanceFromArgs(objectFromArgs *GetAwardsResponse, args map[string]interface{}) *GetAwardsResponse {
	if args["awards"] != nil {
		awardsInterfaceList := args["awards"].([]interface{})
		awards := make([]*Award, 0)

		for _, val := range awardsInterfaceList {
			itemResolved := AwardFromArgs(val.(map[string]interface{}))
			awards = append(awards, itemResolved)
		}
		objectFromArgs.Awards = awards
	}
	return objectFromArgs
}

func (objectFromArgs *GetAwardsResponse) FromArgs(args map[string]interface{}) {
	GetAwardsResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetAwardsResponse) XXX_GraphqlType() *gql.Object {
	return GetAwardsResponseGraphqlType
}

func (msg *GetAwardsResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetAwardsResponseGraphqlArgs
}

func (msg *GetAwardsResponse) XXX_Package() string {
	return "authors"
}

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
		"awards": &gql.Field{
			Type: AwardGraphqlType,
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
		"awards": &gql.InputObjectFieldConfig{
			Type: AwardGraphqlInputType,
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
	"awards": &gql.ArgumentConfig{
		Type: AwardGraphqlInputType,
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
	if args["awards"] != nil {
		val := args["awards"]
		objectFromArgs.Awards = AwardFromArgs(val.(map[string]interface{}))
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

	fields = append(fields, &gql.Field{
		Name: "authors_GetAwards",
		Type: GetAwardsResponseGraphqlType,
		Args: GetAwardsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return AuthorsClientInstance.GetAwards(p.Context, GetAwardsRequestFromArgs(p.Args))
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
