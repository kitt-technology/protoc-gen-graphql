package books

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/grpc"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"context"

	"github.com/graph-gophers/dataloader"

	"github.com/graphql-go/graphql/language/ast"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*graphql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*graphql.Field

var Genre_enum = graphql.NewEnum(graphql.EnumConfig{
	Name: "Genre",
	Values: graphql.EnumValueConfigMap{
		"Biography": &graphql.EnumValueConfig{
			Value: Genre(1),
		},
		"Fiction": &graphql.EnumValueConfig{
			Value: Genre(0),
		},
	},
})

var Genre_type = graphql.NewScalar(graphql.ScalarConfig{
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

var DoNothing_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "DoNothing",
	Fields: graphql.Fields{
		"_null": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var DoNothing_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "DoNothingInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"_null": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})

var DoNothing_args = graphql.FieldConfigArgument{
	"_null": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
}

func DoNothing_from_args(args map[string]interface{}) *DoNothing {
	return DoNothing_instance_from_args(&DoNothing{}, args)
}

func DoNothing_instance_from_args(objectFromArgs *DoNothing, args map[string]interface{}) *DoNothing {
	return objectFromArgs
}

func (objectFromArgs *DoNothing) From_args(args map[string]interface{}) {
	DoNothing_instance_from_args(objectFromArgs, args)
}

func (msg *DoNothing) XXX_type() *graphql.Object {
	return DoNothing_type
}

func (msg *DoNothing) XXX_args() graphql.FieldConfigArgument {
	return DoNothing_args
}

var GetBooksRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "BooksRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
		},
		"hardbackOnly": &graphql.Field{
			Type: graphql.Boolean,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"genres": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(Genre_enum)),
		},
		"releasedAfter": &graphql.Field{
			Type: pg.Timestamp_type,
		},
		"priceGreaterThan": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"copiesGreaterThan": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"copiesLessThan": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"priceLessThan": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"fooBar": &graphql.Field{
			Type: pg.WrappedString,
		},
	},
})

var GetBooksRequest_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BooksRequestInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
		},
		"hardbackOnly": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"genres": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(Genre_enum)),
		},
		"releasedAfter": &graphql.InputObjectFieldConfig{
			Type: pg.Timestamp_input_type,
		},
		"priceGreaterThan": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"copiesGreaterThan": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"copiesLessThan": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"priceLessThan": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"fooBar": &graphql.InputObjectFieldConfig{
			Type: pg.WrappedString,
		},
	},
})

var GetBooksRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(graphql.String)),
	},
	"hardbackOnly": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"price": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"genres": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(Genre_enum)),
	},
	"releasedAfter": &graphql.ArgumentConfig{
		Type: pg.Timestamp_input_type,
	},
	"priceGreaterThan": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
	"copiesGreaterThan": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"copiesLessThan": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"priceLessThan": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
	"fooBar": &graphql.ArgumentConfig{
		Type: pg.WrappedString,
	},
}

func GetBooksRequest_from_args(args map[string]interface{}) *GetBooksRequest {
	return GetBooksRequest_instance_from_args(&GetBooksRequest{}, args)
}

func GetBooksRequest_instance_from_args(objectFromArgs *GetBooksRequest, args map[string]interface{}) *GetBooksRequest {
	if args["ids"] != nil {

		idsInterfaceList := args["ids"].([]interface{})

		var ids []string

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
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = wrapperspb.Float(float32(val.(float64)))
	}
	if args["genres"] != nil {

		genresInterfaceList := args["genres"].([]interface{})

		var genres []Genre

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
	if args["priceGreaterThan"] != nil {
		val := args["priceGreaterThan"]
		objectFromArgs.PriceGreaterThan = float32(val.(float64))
	}
	if args["copiesGreaterThan"] != nil {
		val := args["copiesGreaterThan"]
		objectFromArgs.CopiesGreaterThan = int64(val.(int))
	}
	if args["copiesLessThan"] != nil {
		val := args["copiesLessThan"]
		objectFromArgs.CopiesLessThan = int32(val.(int))
	}
	if args["priceLessThan"] != nil {
		val := args["priceLessThan"]
		objectFromArgs.PriceLessThan = float64(val.(float64))
	}
	if args["fooBar"] != nil {
		val := args["fooBar"]
		objectFromArgs.FooBar = wrapperspb.String(string(val.(string)))
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksRequest) From_args(args map[string]interface{}) {
	GetBooksRequest_instance_from_args(objectFromArgs, args)
}

func (msg *GetBooksRequest) XXX_type() *graphql.Object {
	return GetBooksRequest_type
}

func (msg *GetBooksRequest) XXX_args() graphql.FieldConfigArgument {
	return GetBooksRequest_args
}

var GetBooksResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(Book_type)),
		},
		"foobar": &graphql.Field{
			Type: pg.WrappedString,
		},
	},
})

var GetBooksResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetBooksResponseInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"books": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(Book_input_type)),
		},
		"foobar": &graphql.InputObjectFieldConfig{
			Type: pg.WrappedString,
		},
	},
})

var GetBooksResponse_args = graphql.FieldConfigArgument{
	"books": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(Book_input_type)),
	},
	"foobar": &graphql.ArgumentConfig{
		Type: pg.WrappedString,
	},
}

func GetBooksResponse_from_args(args map[string]interface{}) *GetBooksResponse {
	return GetBooksResponse_instance_from_args(&GetBooksResponse{}, args)
}

func GetBooksResponse_instance_from_args(objectFromArgs *GetBooksResponse, args map[string]interface{}) *GetBooksResponse {
	if args["books"] != nil {

		booksInterfaceList := args["books"].([]interface{})

		var books []*Book

		for _, val := range booksInterfaceList {
			itemResolved := Book_from_args(val.(map[string]interface{}))
			books = append(books, itemResolved)
		}
		objectFromArgs.Books = books
	}
	if args["foobar"] != nil {
		val := args["foobar"]
		objectFromArgs.Foobar = wrapperspb.String(string(val.(string)))
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksResponse) From_args(args map[string]interface{}) {
	GetBooksResponse_instance_from_args(objectFromArgs, args)
}

func (msg *GetBooksResponse) XXX_type() *graphql.Object {
	return GetBooksResponse_type
}

func (msg *GetBooksResponse) XXX_args() graphql.FieldConfigArgument {
	return GetBooksResponse_args
}

var GetBooksByAuthorResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksByAuthorResponse",
	Fields: graphql.Fields{
		"_null": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var GetBooksByAuthorResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetBooksByAuthorResponseInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"_null": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
	},
})

var GetBooksByAuthorResponse_args = graphql.FieldConfigArgument{
	"_null": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
}

func GetBooksByAuthorResponse_from_args(args map[string]interface{}) *GetBooksByAuthorResponse {
	return GetBooksByAuthorResponse_instance_from_args(&GetBooksByAuthorResponse{}, args)
}

func GetBooksByAuthorResponse_instance_from_args(objectFromArgs *GetBooksByAuthorResponse, args map[string]interface{}) *GetBooksByAuthorResponse {
	return objectFromArgs
}

func (objectFromArgs *GetBooksByAuthorResponse) From_args(args map[string]interface{}) {
	GetBooksByAuthorResponse_instance_from_args(objectFromArgs, args)
}

func (msg *GetBooksByAuthorResponse) XXX_type() *graphql.Object {
	return GetBooksByAuthorResponse_type
}

func (msg *GetBooksByAuthorResponse) XXX_args() graphql.FieldConfigArgument {
	return GetBooksByAuthorResponse_args
}

var BooksByAuthor_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "BooksByAuthor",
	Fields: graphql.Fields{
		"results": &graphql.Field{
			Type: graphql.NewList(graphql.NewNonNull(Book_type)),
		},
	},
})

var BooksByAuthor_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BooksByAuthorInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"results": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.NewNonNull(Book_input_type)),
		},
	},
})

var BooksByAuthor_args = graphql.FieldConfigArgument{
	"results": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.NewNonNull(Book_input_type)),
	},
}

func BooksByAuthor_from_args(args map[string]interface{}) *BooksByAuthor {
	return BooksByAuthor_instance_from_args(&BooksByAuthor{}, args)
}

func BooksByAuthor_instance_from_args(objectFromArgs *BooksByAuthor, args map[string]interface{}) *BooksByAuthor {
	if args["results"] != nil {

		resultsInterfaceList := args["results"].([]interface{})

		var results []*Book

		for _, val := range resultsInterfaceList {
			itemResolved := Book_from_args(val.(map[string]interface{}))
			results = append(results, itemResolved)
		}
		objectFromArgs.Results = results
	}
	return objectFromArgs
}

func (objectFromArgs *BooksByAuthor) From_args(args map[string]interface{}) {
	BooksByAuthor_instance_from_args(objectFromArgs, args)
}

func (msg *BooksByAuthor) XXX_type() *graphql.Object {
	return BooksByAuthor_type
}

func (msg *BooksByAuthor) XXX_args() graphql.FieldConfigArgument {
	return BooksByAuthor_args
}

var Book_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"authorId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"genre": &graphql.Field{
			Type: Genre_enum,
		},
		"releaseDate": &graphql.Field{
			Type: pg.Timestamp_type,
		},
		"price": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"copies": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var Book_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BookInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"authorId": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"genre": &graphql.InputObjectFieldConfig{
			Type: Genre_enum,
		},
		"releaseDate": &graphql.InputObjectFieldConfig{
			Type: pg.Timestamp_input_type,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"copies": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var Book_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"authorId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"genre": &graphql.ArgumentConfig{
		Type: Genre_enum,
	},
	"releaseDate": &graphql.ArgumentConfig{
		Type: pg.Timestamp_input_type,
	},
	"price": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Float),
	},
	"copies": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

func Book_from_args(args map[string]interface{}) *Book {
	return Book_instance_from_args(&Book{}, args)
}

func Book_instance_from_args(objectFromArgs *Book, args map[string]interface{}) *Book {
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
		objectFromArgs.Copies = int64(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *Book) From_args(args map[string]interface{}) {
	Book_instance_from_args(objectFromArgs, args)
}

func (msg *Book) XXX_type() *graphql.Object {
	return Book_type
}

func (msg *Book) XXX_args() graphql.FieldConfigArgument {
	return Book_args
}

var BooksClientInstance BooksClient

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		BooksClientInstance = NewBooksClient(pg.GrpcConnection("localhost:50051", opts...))
	})
	fields = append(fields, &graphql.Field{
		Name: "Books_GetBooks",
		Type: GetBooksResponse_type,
		Args: GetBooksRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return BooksClientInstance.GetBooks(p.Context, GetBooksRequest_from_args(p.Args))
		},
	})

	fields = append(fields, &graphql.Field{
		Name: "Books_DoNothing",
		Type: DoNothing_type,
		Args: DoNothing_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return BooksClientInstance.DoNothing(p.Context, DoNothing_from_args(p.Args))
		},
	})

}

func WithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "BooksByAuthorLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			resp, err := BooksClientInstance.GetBooksByAuthor(ctx, &pg.BatchRequest{
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

func LoadBooksByAuthor(p graphql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("BooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("BooksByAuthorLoader").(*dataloader.Loader)
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

func LoadManyBooksByAuthor(p graphql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("BooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("BooksByAuthorLoader").(*dataloader.Loader)
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
