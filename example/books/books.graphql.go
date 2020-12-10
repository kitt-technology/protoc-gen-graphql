package books

import (
	"github.com/graphql-go/graphql"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"github.com/graph-gophers/dataloader"
	"context"
)

var Fields []*graphql.Field

var GetBooksRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var GetBooksRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
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

	return &objectFromArgs
}

var GetBooksResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(Book_type),
		},
	},
})

var GetBooksResponse_args = graphql.FieldConfigArgument{
	"books": &graphql.ArgumentConfig{
		Type: graphql.NewList(Book_type),
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

var GetBooksByAuthorResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksByAuthorResponse",
	Fields: graphql.Fields{
		"results": &graphql.Field{
			Type: graphql.NewList(BooksByAuthor_type),
		},
	},
})

var GetBooksByAuthorResponse_args = graphql.FieldConfigArgument{
	"results": &graphql.ArgumentConfig{
		Type: graphql.NewList(BooksByAuthor_type),
	},
}

func GetBooksByAuthorResponse_from_args(args map[string]interface{}) *GetBooksByAuthorResponse {
	objectFromArgs := GetBooksByAuthorResponse{}
	if args["results"] != nil {

		resultsInterfaceList := args["results"].([]interface{})

		var results []*BooksByAuthor
		for _, item := range resultsInterfaceList {
			results = append(results, item.(*BooksByAuthor))
		}
		objectFromArgs.Results = results

	}

	return &objectFromArgs
}

var BooksByAuthor_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "BooksByAuthor",
	Fields: graphql.Fields{
		"authorId": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"books": &graphql.Field{
			Type: graphql.NewList(Book_type),
		},
	},
})

var BooksByAuthor_args = graphql.FieldConfigArgument{
	"authorId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"books": &graphql.ArgumentConfig{
		Type: graphql.NewList(Book_type),
	},
}

func BooksByAuthor_from_args(args map[string]interface{}) *BooksByAuthor {
	objectFromArgs := BooksByAuthor{}
	objectFromArgs.AuthorId = args["authorId"].(string)
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
}

func Book_from_args(args map[string]interface{}) *Book {
	objectFromArgs := Book{}
	objectFromArgs.Id = args["id"].(string)
	objectFromArgs.Name = args["name"].(string)
	objectFromArgs.AuthorId = args["authorId"].(string)

	return &objectFromArgs
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

	Fields = append(Fields, &graphql.Field{
		Name: "Books_GetBooksByAuthor",
		Type: GetBooksByAuthorResponse_type,
		Args: GetBooksRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return client.GetBooksByAuthor(p.Context, GetBooksRequest_from_args(p.Args))
		},
	})

}

func LoadBooksByAuthor(originalContext context.Context, key string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.GetBooksByAuthor(ctx, &GetBooksRequest{
			Ids: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, item := range resp.Results {
			results = append(results, &dataloader.Result{Data: item})
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

		resp, err := client.GetBooksByAuthor(ctx, &GetBooksRequest{
			Ids: keys.Keys(),
		})

		if err != nil {
			return results
		}

		for _, item := range resp.Results {
			results = append(results, &dataloader.Result{Data: item})
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
