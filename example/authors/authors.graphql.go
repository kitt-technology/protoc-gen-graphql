package authors

import (
	"github.com/graphql-go/graphql"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"github.com/graph-gophers/dataloader"
	"context"
)

var Fields []*graphql.Field

var GetAuthorsRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var GetAuthorsRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
}

func GetAuthorsRequest_from_args(args map[string]interface{}) *GetAuthorsRequest {
	objectFromArgs := GetAuthorsRequest{}
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

var GetAuthorsResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsResponse",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewList(Author_type),
		},
	},
})

var GetAuthorsResponse_args = graphql.FieldConfigArgument{
	"authors": &graphql.ArgumentConfig{
		Type: graphql.NewList(Author_type),
	},
}

func GetAuthorsResponse_from_args(args map[string]interface{}) *GetAuthorsResponse {
	objectFromArgs := GetAuthorsResponse{}
	if args["authors"] != nil {

		authorsInterfaceList := args["authors"].([]interface{})

		var authors []*Author
		for _, item := range authorsInterfaceList {
			authors = append(authors, item.(*Author))
		}
		objectFromArgs.Authors = authors

	}

	return &objectFromArgs
}

var AuthorsBatchRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthorsBatchRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var AuthorsBatchRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.String),
	},
}

func AuthorsBatchRequest_from_args(args map[string]interface{}) *AuthorsBatchRequest {
	objectFromArgs := AuthorsBatchRequest{}
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

var AuthorsBatchResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthorsBatchResponse",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AuthorsBatchResponse_args = graphql.FieldConfigArgument{
	"message": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func AuthorsBatchResponse_from_args(args map[string]interface{}) *AuthorsBatchResponse {
	objectFromArgs := AuthorsBatchResponse{}

	return &objectFromArgs
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

var Author_args = graphql.FieldConfigArgument{
	"id": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

func Author_from_args(args map[string]interface{}) *Author {
	objectFromArgs := Author{}

	objectFromArgs.Id = args["id"].(string)

	objectFromArgs.Name = args["name"].(string)

	return &objectFromArgs
}

var client AuthorsClient

func init() {
	client = NewAuthorsClient(pg.GrpcConnection("localhost:50052"))
	Fields = append(Fields, &graphql.Field{
		Name: "Authors_GetAuthors",
		Type: GetAuthorsResponse_type,
		Args: GetAuthorsRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return client.GetAuthors(p.Context, GetAuthorsRequest_from_args(p.Args))
		},
	})

}

func LoadAuthor(originalContext context.Context, key string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.LoadAuthors(ctx, &pg.BatchRequest{
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
		return res.(*Author), nil
	}, nil
}

func LoadManyAuthor(originalContext context.Context, keys []string) (func() (interface{}, error), error) {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		resp, err := client.LoadAuthors(ctx, &pg.BatchRequest{
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

		var results []*Author
		for _, res := range resSlice {
			results = append(results, res.(*Author))
		}

		return results, nil
	}, nil
}
