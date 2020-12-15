package cases

import (
	"github.com/graphql-go/graphql"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"context"
	"github.com/graph-gophers/dataloader"
)

var Fields []*graphql.Field

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

var GetBooksRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "BooksRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"hardbackOnly": &graphql.Field{
			Type: graphql.Boolean,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"genres": &graphql.Field{
			Type: graphql.NewList(Genre_enum),
		},
		"releasedAfter": &graphql.Field{
			Type: pg.Timestamp_type,
		},
		"pagination": &graphql.Field{
			Type: PaginationOptions_type,
		},
		"filters": &graphql.Field{
			Type: graphql.NewList(Filter_type),
		},
	},
})

var GetBooksRequest_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BooksRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"ids": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
		"hardbackOnly": &graphql.InputObjectFieldConfig{
			Type: graphql.Boolean,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"genres": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(Genre_enum),
		},
		"releasedAfter": &graphql.InputObjectFieldConfig{
			Type: pg.Timestamp_input_type,
		},
		"pagination": &graphql.InputObjectFieldConfig{
			Type: PaginationOptions_input_type,
		},
		"filters": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(Filter_input_type),
		},
	},
})

var GetBooksRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
	"hardbackOnly": &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	},
	"price": &graphql.ArgumentConfig{
		Type: graphql.Float,
	},
	"genres": &graphql.ArgumentConfig{
		Type: graphql.NewList(Genre_enum),
	},
	"releasedAfter": &graphql.ArgumentConfig{
		Type: pg.Timestamp_input_type,
	},
	"pagination": &graphql.ArgumentConfig{
		Type: PaginationOptions_input_type,
	},
	"filters": &graphql.ArgumentConfig{
		Type: graphql.NewList(Filter_input_type),
	},
}

func GetBooksRequest_from_args(args map[string]interface{}) *GetBooksRequest {
	objectFromArgs := GetBooksRequest{}
	if args["ids"] != nil {

		idsInterfaceList := args["ids"].([]interface{})

		var ids []string
		for _, item := range idsInterfaceList {
			ids = append(ids, item.(string))
		}
		objectFromArgs.Ids = ids

	}

	if args["hardbackOnly"] != nil {
		objectFromArgs.HardbackOnly = wrapperspb.Bool(args["hardbackOnly"].(bool))
	}

	if args["price"] != nil {
		objectFromArgs.Price = wrapperspb.Float(args["price"].(float32))
	}

	if args["genres"] != nil {

		genresInterfaceList := args["genres"].([]interface{})

		var genres []Genre
		for _, item := range genresInterfaceList {
			genres = append(genres, item.(Genre))
		}
		objectFromArgs.Genres = genres

	}

	if args["pagination"] != nil {
		objectFromArgs.Pagination = PaginationOptions_from_args(args["pagination"].(map[string]interface{}))
	}

	if args["filters"] != nil {

		filtersInterfaceList := args["filters"].([]interface{})

		var filters []*Filter
		for _, item := range filtersInterfaceList {
			filters = append(filters, item.(*Filter))
		}
		objectFromArgs.Filters = filters

	}

	return &objectFromArgs
}

func (objectFromArgs *GetBooksRequest) From_args(args map[string]interface{}) {
	objectFromArgs = GetBooksRequest_from_args(args)

}

func (msg *GetBooksRequest) XXX_type() *graphql.Object {
	return GetBooksRequest_type
}

func (msg *GetBooksRequest) XXX_args() graphql.FieldConfigArgument {
	return GetBooksRequest_args
}

var PaginationOptions_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaginationOptions",
	Fields: graphql.Fields{
		"page": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"perPage": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var PaginationOptions_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaginationOptions",
	Fields: graphql.InputObjectConfigFieldMap{
		"page": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"perPage": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

var PaginationOptions_args = graphql.FieldConfigArgument{
	"page": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
	"perPage": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	},
}

func PaginationOptions_from_args(args map[string]interface{}) *PaginationOptions {
	objectFromArgs := PaginationOptions{}

	objectFromArgs.Page = args["page"].(int)

	objectFromArgs.PerPage = args["perPage"].(int)

	return &objectFromArgs
}

func (objectFromArgs *PaginationOptions) From_args(args map[string]interface{}) {
	objectFromArgs = PaginationOptions_from_args(args)

}

func (msg *PaginationOptions) XXX_type() *graphql.Object {
	return PaginationOptions_type
}

func (msg *PaginationOptions) XXX_args() graphql.FieldConfigArgument {
	return PaginationOptions_args
}

var Filter_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "Filter",
	Fields: graphql.Fields{
		"query": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var Filter_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Filter",
	Fields: graphql.InputObjectConfigFieldMap{
		"query": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var Filter_args = graphql.FieldConfigArgument{
	"query": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func Filter_from_args(args map[string]interface{}) *Filter {
	objectFromArgs := Filter{}

	objectFromArgs.Query = args["query"].(string)

	return &objectFromArgs
}

func (objectFromArgs *Filter) From_args(args map[string]interface{}) {
	objectFromArgs = Filter_from_args(args)

}

func (msg *Filter) XXX_type() *graphql.Object {
	return Filter_type
}

func (msg *Filter) XXX_args() graphql.FieldConfigArgument {
	return Filter_args
}

var GetBooksResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(Book_type),
		},
	},
})

var GetBooksResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetBooksResponse",
	Fields: graphql.InputObjectConfigFieldMap{
		"books": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(Book_input_type),
		},
	},
})

var GetBooksResponse_args = graphql.FieldConfigArgument{
	"books": &graphql.ArgumentConfig{
		Type: graphql.NewList(Book_input_type),
	},
}

func GetBooksResponse_from_args(args map[string]interface{}) *GetBooksResponse {
	objectFromArgs := GetBooksResponse{}
	if args["books"] != nil {

		booksInterfaceList := args["books"].([]interface{})

		var books []*Book
		for _, item := range booksInterfaceList {
			books = append(books, item.(*Book))
		}
		objectFromArgs.Books = books

	}

	return &objectFromArgs
}

func (objectFromArgs *GetBooksResponse) From_args(args map[string]interface{}) {
	objectFromArgs = GetBooksResponse_from_args(args)

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
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var GetBooksByAuthorResponse_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "GetBooksByAuthorResponse",
	Fields: graphql.InputObjectConfigFieldMap{
		"message": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var GetBooksByAuthorResponse_args = graphql.FieldConfigArgument{
	"message": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func GetBooksByAuthorResponse_from_args(args map[string]interface{}) *GetBooksByAuthorResponse {
	objectFromArgs := GetBooksByAuthorResponse{}

	return &objectFromArgs
}

func (objectFromArgs *GetBooksByAuthorResponse) From_args(args map[string]interface{}) {
	objectFromArgs = GetBooksByAuthorResponse_from_args(args)

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
			Type: graphql.NewList(Book_type),
		},
	},
})

var BooksByAuthor_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BooksByAuthor",
	Fields: graphql.InputObjectConfigFieldMap{
		"results": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(Book_input_type),
		},
	},
})

var BooksByAuthor_args = graphql.FieldConfigArgument{
	"results": &graphql.ArgumentConfig{
		Type: graphql.NewList(Book_input_type),
	},
}

func BooksByAuthor_from_args(args map[string]interface{}) *BooksByAuthor {
	objectFromArgs := BooksByAuthor{}
	if args["results"] != nil {

		resultsInterfaceList := args["results"].([]interface{})

		var results []*Book
		for _, item := range resultsInterfaceList {
			results = append(results, item.(*Book))
		}
		objectFromArgs.Results = results

	}

	return &objectFromArgs
}

func (objectFromArgs *BooksByAuthor) From_args(args map[string]interface{}) {
	objectFromArgs = BooksByAuthor_from_args(args)

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
	},
})

var Book_input_type = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Book",
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
}

func Book_from_args(args map[string]interface{}) *Book {
	objectFromArgs := Book{}

	objectFromArgs.Id = args["id"].(string)

	objectFromArgs.Name = args["name"].(string)

	objectFromArgs.AuthorId = args["authorId"].(string)

	objectFromArgs.Genre = args["genre"].(Genre)

	return &objectFromArgs
}

func (objectFromArgs *Book) From_args(args map[string]interface{}) {
	objectFromArgs = Book_from_args(args)

}

func (msg *Book) XXX_type() *graphql.Object {
	return Book_type
}

func (msg *Book) XXX_args() graphql.FieldConfigArgument {
	return Book_args
}

var client BooksClient

func init() {
	client = NewBooksClient(pg.GrpcConnection("localhost:50051"))
	Fields = append(Fields, &graphql.Field{
		Name: "Books_GetBooks",
		Type: GetBooksResponse_type,
		Args: GetBooksRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return client.GetBooks(p.Context, GetBooksRequest_from_args(p.Args))
		},
	})

}

func LoadBooksByAuthor(originalContext context.Context, key string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.GetBooksByAuthor(ctx, &pg.BatchRequest{
			Keys: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, key := range keys.Keys() {
			results = append(results, &dataloader.Result{Data: resp.Results[key]})
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	thunk := loader.Load(originalContext, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*BooksByAuthor), nil
	}, nil
}

func LoadManyBooksByAuthor(originalContext context.Context, keys []string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.GetBooksByAuthor(ctx, &pg.BatchRequest{
			Keys: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, key := range keys.Keys() {
			results = append(results, &dataloader.Result{Data: resp.Results[key]})
		}

		return results
	}

	loader := dataloader.NewBatchedLoader(batchFn)

	thunk := loader.LoadMany(originalContext, dataloader.NewKeysFromStrings(keys))
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
