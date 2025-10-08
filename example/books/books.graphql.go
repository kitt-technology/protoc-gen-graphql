package books

import (
	"context"
	"os"

	"github.com/graph-gophers/dataloader"
	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	common_example "github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var GenreGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "Genre",
	Values: gql.EnumValueConfigMap{
		"Biography": &gql.EnumValueConfig{
			Value: Genre(1),
		},
		"Fiction": &gql.EnumValueConfig{
			Value: Genre(0),
		},
	},
})

var GenreGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "Genre",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(Genre).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var DoNothingGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "DoNothing",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var DoNothingGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "DoNothingInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var DoNothingGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func DoNothingFromArgs(args map[string]interface{}) *DoNothing {
	return DoNothingInstanceFromArgs(&DoNothing{}, args)
}

func DoNothingInstanceFromArgs(objectFromArgs *DoNothing, args map[string]interface{}) *DoNothing {
	return objectFromArgs
}

func (objectFromArgs *DoNothing) FromArgs(args map[string]interface{}) {
	DoNothingInstanceFromArgs(objectFromArgs, args)
}

func (msg *DoNothing) XXX_GraphqlType() *gql.Object {
	return DoNothingGraphqlType
}

func (msg *DoNothing) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return DoNothingGraphqlArgs
}

func (msg *DoNothing) XXX_Package() string {
	return "books"
}

var GetBooksRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "BooksRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"hardbackOnly": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*GetBooksRequest) == nil || p.Source.(*GetBooksRequest).HardbackOnly == nil {
					return nil, nil
				}
				return p.Source.(*GetBooksRequest).HardbackOnly.Value, nil
			},
		},
		"genres": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
		},
		"releasedAfter": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
	},
})

var GetBooksRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BooksRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"hardbackOnly": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"genres": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
		},
		"releasedAfter": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
	},
})

var GetBooksRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(gql.String)),
	},
	"hardbackOnly": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"genres": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
	},
	"releasedAfter": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
}

func GetBooksRequestFromArgs(args map[string]interface{}) *GetBooksRequest {
	return GetBooksRequestInstanceFromArgs(&GetBooksRequest{}, args)
}

func GetBooksRequestInstanceFromArgs(objectFromArgs *GetBooksRequest, args map[string]interface{}) *GetBooksRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		ids := make([]string, 0)

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	if args["hardbackOnly"] != nil {
		val := args["hardbackOnly"]
		objectFromArgs.HardbackOnly = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["genres"] != nil {
		genresInterfaceList := args["genres"].([]interface{})
		genres := make([]Genre, 0)

		for _, val := range genresInterfaceList {
			itemResolved := val.(Genre)
			genres = append(genres, itemResolved)
		}
		objectFromArgs.Genres = genres
	}
	if args["releasedAfter"] != nil {
		val := args["releasedAfter"]
		objectFromArgs.ReleasedAfter = pg.ToTimestamp(val)
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksRequest) FromArgs(args map[string]interface{}) {
	GetBooksRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksRequest) XXX_GraphqlType() *gql.Object {
	return GetBooksRequestGraphqlType
}

func (msg *GetBooksRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksRequestGraphqlArgs
}

func (msg *GetBooksRequest) XXX_Package() string {
	return "books"
}

var GetBooksResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: gql.Fields{
		"books": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlType)),
		},
	},
})

var GetBooksResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"books": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
		},
	},
})

var GetBooksResponseGraphqlArgs = gql.FieldConfigArgument{
	"books": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
	},
}

func GetBooksResponseFromArgs(args map[string]interface{}) *GetBooksResponse {
	return GetBooksResponseInstanceFromArgs(&GetBooksResponse{}, args)
}

func GetBooksResponseInstanceFromArgs(objectFromArgs *GetBooksResponse, args map[string]interface{}) *GetBooksResponse {
	if args["books"] != nil {
		booksInterfaceList := args["books"].([]interface{})
		books := make([]*Book, 0)

		for _, val := range booksInterfaceList {
			itemResolved := BookFromArgs(val.(map[string]interface{}))
			books = append(books, itemResolved)
		}
		objectFromArgs.Books = books
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksResponse) FromArgs(args map[string]interface{}) {
	GetBooksResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksResponseGraphqlType
}

func (msg *GetBooksResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksResponseGraphqlArgs
}

func (msg *GetBooksResponse) XXX_Package() string {
	return "books"
}

var GetBooksByAuthorRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksByAuthorRequest",
	Fields: gql.Fields{
		"keys": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var GetBooksByAuthorRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksByAuthorRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"keys": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var GetBooksByAuthorRequestGraphqlArgs = gql.FieldConfigArgument{
	"keys": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func GetBooksByAuthorRequestFromArgs(args map[string]interface{}) *GetBooksByAuthorRequest {
	return GetBooksByAuthorRequestInstanceFromArgs(&GetBooksByAuthorRequest{}, args)
}

func GetBooksByAuthorRequestInstanceFromArgs(objectFromArgs *GetBooksByAuthorRequest, args map[string]interface{}) *GetBooksByAuthorRequest {
	if args["keys"] != nil {
		keysInterfaceList := args["keys"].([]interface{})
		keys := make([]string, 0)

		for _, val := range keysInterfaceList {
			itemResolved := string(val.(string))
			keys = append(keys, itemResolved)
		}
		objectFromArgs.Keys = keys
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksByAuthorRequest) FromArgs(args map[string]interface{}) {
	GetBooksByAuthorRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksByAuthorRequest) XXX_GraphqlType() *gql.Object {
	return GetBooksByAuthorRequestGraphqlType
}

func (msg *GetBooksByAuthorRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksByAuthorRequestGraphqlArgs
}

func (msg *GetBooksByAuthorRequest) XXX_Package() string {
	return "books"
}

var GetBooksByAuthorResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksByAuthorResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var GetBooksByAuthorResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksByAuthorResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetBooksByAuthorResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetBooksByAuthorResponseFromArgs(args map[string]interface{}) *GetBooksByAuthorResponse {
	return GetBooksByAuthorResponseInstanceFromArgs(&GetBooksByAuthorResponse{}, args)
}

func GetBooksByAuthorResponseInstanceFromArgs(objectFromArgs *GetBooksByAuthorResponse, args map[string]interface{}) *GetBooksByAuthorResponse {
	return objectFromArgs
}

func (objectFromArgs *GetBooksByAuthorResponse) FromArgs(args map[string]interface{}) {
	GetBooksByAuthorResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksByAuthorResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksByAuthorResponseGraphqlType
}

func (msg *GetBooksByAuthorResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksByAuthorResponseGraphqlArgs
}

func (msg *GetBooksByAuthorResponse) XXX_Package() string {
	return "books"
}

var GetBooksBatchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksBatchRequest",
	Fields: gql.Fields{
		"reqs": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlType)),
		},
	},
})

var GetBooksBatchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksBatchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"reqs": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlInputType)),
		},
	},
})

var GetBooksBatchRequestGraphqlArgs = gql.FieldConfigArgument{
	"reqs": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlInputType)),
	},
}

func GetBooksBatchRequestFromArgs(args map[string]interface{}) *GetBooksBatchRequest {
	return GetBooksBatchRequestInstanceFromArgs(&GetBooksBatchRequest{}, args)
}

func GetBooksBatchRequestInstanceFromArgs(objectFromArgs *GetBooksBatchRequest, args map[string]interface{}) *GetBooksBatchRequest {
	if args["reqs"] != nil {
		reqsInterfaceList := args["reqs"].([]interface{})
		reqs := make([]*GetBooksRequest, 0)

		for _, val := range reqsInterfaceList {
			itemResolved := GetBooksRequestFromArgs(val.(map[string]interface{}))
			reqs = append(reqs, itemResolved)
		}
		objectFromArgs.Reqs = reqs
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksBatchRequest) FromArgs(args map[string]interface{}) {
	GetBooksBatchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksBatchRequest) XXX_GraphqlType() *gql.Object {
	return GetBooksBatchRequestGraphqlType
}

func (msg *GetBooksBatchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksBatchRequestGraphqlArgs
}

func (msg *GetBooksBatchRequest) XXX_Package() string {
	return "books"
}

var GetBooksBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var GetBooksBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetBooksBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetBooksBatchResponseFromArgs(args map[string]interface{}) *GetBooksBatchResponse {
	return GetBooksBatchResponseInstanceFromArgs(&GetBooksBatchResponse{}, args)
}

func GetBooksBatchResponseInstanceFromArgs(objectFromArgs *GetBooksBatchResponse, args map[string]interface{}) *GetBooksBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *GetBooksBatchResponse) FromArgs(args map[string]interface{}) {
	GetBooksBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksBatchResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksBatchResponseGraphqlType
}

func (msg *GetBooksBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksBatchResponseGraphqlArgs
}

func (msg *GetBooksBatchResponse) XXX_Package() string {
	return "books"
}

var BooksByAuthorGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "BooksByAuthor",
	Fields: gql.Fields{
		"results": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlType)),
		},
	},
})

var BooksByAuthorGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BooksByAuthorInput",
	Fields: gql.InputObjectConfigFieldMap{
		"results": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
		},
	},
})

var BooksByAuthorGraphqlArgs = gql.FieldConfigArgument{
	"results": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
	},
}

func BooksByAuthorFromArgs(args map[string]interface{}) *BooksByAuthor {
	return BooksByAuthorInstanceFromArgs(&BooksByAuthor{}, args)
}

func BooksByAuthorInstanceFromArgs(objectFromArgs *BooksByAuthor, args map[string]interface{}) *BooksByAuthor {
	if args["results"] != nil {
		resultsInterfaceList := args["results"].([]interface{})
		results := make([]*Book, 0)

		for _, val := range resultsInterfaceList {
			itemResolved := BookFromArgs(val.(map[string]interface{}))
			results = append(results, itemResolved)
		}
		objectFromArgs.Results = results
	}
	return objectFromArgs
}

func (objectFromArgs *BooksByAuthor) FromArgs(args map[string]interface{}) {
	BooksByAuthorInstanceFromArgs(objectFromArgs, args)
}

func (msg *BooksByAuthor) XXX_GraphqlType() *gql.Object {
	return BooksByAuthorGraphqlType
}

func (msg *BooksByAuthor) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return BooksByAuthorGraphqlArgs
}

func (msg *BooksByAuthor) XXX_Package() string {
	return "books"
}

var BookGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Book",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"authorId": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"genre": &gql.Field{
			Type: GenreGraphqlEnum,
		},
		"releaseDate": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"price": &gql.Field{
			Type: gql.NewNonNull(gql.Float),
		},
		"copies": &gql.Field{
			Type: gql.Int,
		},
		"priceTwo": &gql.Field{
			Type: common_example.MoneyGraphqlType,
		},
		"isSigned": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*Book) == nil || p.Source.(*Book).IsSigned == nil {
					return nil, nil
				}
				return p.Source.(*Book).IsSigned.Value, nil
			},
		},
	},
})

var BookGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BookInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"authorId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"genre": &gql.InputObjectFieldConfig{
			Type: GenreGraphqlEnum,
		},
		"releaseDate": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"price": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Float),
		},
		"copies": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
		"priceTwo": &gql.InputObjectFieldConfig{
			Type: common_example.MoneyGraphqlInputType,
		},
		"isSigned": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var BookGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"authorId": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"genre": &gql.ArgumentConfig{
		Type: GenreGraphqlEnum,
	},
	"releaseDate": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"price": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Float),
	},
	"copies": &gql.ArgumentConfig{
		Type: gql.Int,
	},
	"priceTwo": &gql.ArgumentConfig{
		Type: common_example.MoneyGraphqlInputType,
	},
	"isSigned": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func BookFromArgs(args map[string]interface{}) *Book {
	return BookInstanceFromArgs(&Book{}, args)
}

func BookInstanceFromArgs(objectFromArgs *Book, args map[string]interface{}) *Book {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["name"] != nil {
		val := args["name"]
		objectFromArgs.Name = string(val.(string))
	}
	if args["authorId"] != nil {
		val := args["authorId"]
		objectFromArgs.AuthorId = string(val.(string))
	}
	if args["genre"] != nil {
		val := args["genre"]
		objectFromArgs.Genre = val.(Genre)
	}
	if args["releaseDate"] != nil {
		val := args["releaseDate"]
		objectFromArgs.ReleaseDate = pg.ToTimestamp(val)
	}
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = float32(val.(float64))
	}
	if args["copies"] != nil {
		val := args["copies"]
		ptr := int64(val.(int))
		objectFromArgs.Copies = &ptr
	}
	if args["priceTwo"] != nil {
		val := args["priceTwo"]
		objectFromArgs.PriceTwo = common_example.MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["isSigned"] != nil {
		val := args["isSigned"]
		objectFromArgs.IsSigned = wrapperspb.Bool(bool(val.(bool)))
	}
	return objectFromArgs
}

func (objectFromArgs *Book) FromArgs(args map[string]interface{}) {
	BookInstanceFromArgs(objectFromArgs, args)
}

func (msg *Book) XXX_GraphqlType() *gql.Object {
	return BookGraphqlType
}

func (msg *Book) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return BookGraphqlArgs
}

func (msg *Book) XXX_Package() string {
	return "books"
}

var BooksClientInstance BooksClient
var BooksServiceInstance BooksServer
var BooksDialOpts []grpc.DialOption

type BooksOption func(*BooksConfig)

type BooksConfig struct {
	service  BooksServer
	client   BooksClient
	dialOpts []grpc.DialOption
}

// WithService sets the service implementation for direct calls (no gRPC)
func WithService(service BooksServer) BooksOption {
	return func(cfg *BooksConfig) {
		cfg.service = service
	}
}

// WithClient sets the gRPC client for remote calls
func WithClient(client BooksClient) BooksOption {
	return func(cfg *BooksConfig) {
		cfg.client = client
	}
}

// WithDialOptions sets the dial options for the gRPC client
func WithDialOptions(opts ...grpc.DialOption) BooksOption {
	return func(cfg *BooksConfig) {
		cfg.dialOpts = opts
	}
}

func Init(ctx context.Context, opts ...BooksOption) (context.Context, []*gql.Field) {
	cfg := &BooksConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	BooksServiceInstance = cfg.service
	BooksClientInstance = cfg.client
	BooksDialOpts = cfg.dialOpts

	var fields []*gql.Field
	fields = append(fields, &gql.Field{
		Name: "books_GetBooks",
		Type: GetBooksResponseGraphqlType,
		Args: GetBooksRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if BooksServiceInstance != nil {
				return BooksServiceInstance.GetBooks(p.Context, GetBooksRequestFromArgs(p.Args))
			}
			if BooksClientInstance == nil {
				BooksClientInstance = getBooksClient()
			}
			return BooksClientInstance.GetBooks(p.Context, GetBooksRequestFromArgs(p.Args))
		},
	})

	ctx = BooksWithLoaders(ctx)

	return ctx, fields
}

func getBooksClient() BooksClient {
	host := "localhost:50051"
	envHost := os.Getenv("SERVICE_HOST")
	if envHost != "" {
		host = envHost
	}
	return NewBooksClient(pg.GrpcConnection(host, BooksDialOpts...))
}

// SetBooksService sets the service implementation for direct calls (no gRPC)
func SetBooksService(service BooksServer) {
	BooksServiceInstance = service
}

// SetBooksClient sets the gRPC client for remote calls
func SetBooksClient(client BooksClient) {
	BooksClientInstance = client
}

func BooksWithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "GetBooksByAuthorLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *GetBooksByAuthorResponse
			var err error
			if BooksServiceInstance != nil {
				resp, err = BooksServiceInstance.GetBooksByAuthor(ctx, &GetBooksByAuthorRequest{
					Keys: keys.Keys(),
				})
			} else {
				if BooksClientInstance == nil {
					BooksClientInstance = getBooksClient()
				}
				resp, err = BooksClientInstance.GetBooksByAuthor(ctx, &GetBooksByAuthorRequest{
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
					var empty *BooksByAuthor
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	ctx = context.WithValue(ctx, "GetBooksBatchLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var requests []*GetBooksRequest
			for _, key := range keys {
				requests = append(requests, key.(*GetBooksRequestKey).GetBooksRequest)
			}
			var resp *GetBooksBatchResponse
			var err error
			if BooksServiceInstance != nil {
				resp, err = BooksServiceInstance.GetBooksBatch(ctx, &GetBooksBatchRequest{
					Reqs: requests,
				})
			} else {
				if BooksClientInstance == nil {
					BooksClientInstance = getBooksClient()
				}
				resp, err = BooksClientInstance.GetBooksBatch(ctx, &GetBooksBatchRequest{
					Reqs: requests,
				})
			}

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *GetBooksResponse
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	return ctx
}

func GetBooksByAuthor(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksByAuthorLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*BooksByAuthor), nil
	}, nil
}

func GetBooksByAuthorMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksByAuthorLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))

	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*BooksByAuthor
		for _, res := range resSlice {
			results = append(results, res.(*BooksByAuthor))
		}

		return results, nil
	}, nil
}

type GetBooksRequestKey struct {
	*GetBooksRequest
}

func (key *GetBooksRequestKey) String() string {
	return pg.ProtoKey(key)
}

func (key *GetBooksRequestKey) Raw() interface{} {
	return key
}

func GetBooksBatch(p gql.ResolveParams, key *GetBooksRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksBatchLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksBatchLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, &GetBooksRequestKey{key})
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*GetBooksResponse), nil
	}, nil
}

func GetBooksBatchMany(p gql.ResolveParams, keys []*GetBooksRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksBatchLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksBatchLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	loaderKeys := make(dataloader.Keys, len(keys))
	for ix := range keys {
		loaderKeys[ix] = &GetBooksRequestKey{keys[ix]}
	}

	thunk := loader.LoadMany(p.Context, loaderKeys)

	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*GetBooksResponse
		for _, res := range resSlice {
			results = append(results, res.(*GetBooksResponse))
		}

		return results, nil
	}, nil
}
