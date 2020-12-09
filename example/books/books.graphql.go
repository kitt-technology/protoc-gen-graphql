package books

import "github.com/graphql-go/graphql"
import pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
import "google.golang.org/protobuf/proto"

var mutations []*graphql.Field
var queries []*graphql.Field

var mutationResolver func(command proto.Message, success proto.Message) (proto.Message, error)
var dataloadersToRegister map[string][]pg.RegisterDataloaderFn
var dataloadersToProvide map[string]pg.Dataloader

func Register(config pg.ProtoConfig) pg.ProtoConfig {
	config.Queries = append(config.Queries, queries...)
	return config
}

var GetBooksRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksRequest",
	Fields: graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var GetBooksRequest_args = graphql.FieldConfigArgument{
	"message": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func GetBooksRequest_from_args(args map[string]interface{}) *GetBooksRequest {
	return &GetBooksRequest{}
}

var GetBooksResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(Book_type)),
		},
	},
})

var GetBooksResponse_args = graphql.FieldConfigArgument{
	"books": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.NewList(Book_type)),
	},
}

func GetBooksResponse_from_args(args map[string]interface{}) *GetBooksResponse {
	return &GetBooksResponse{
		Books: args["books"].([]*Book),
	}
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
	return &Book{
		Id:       args["id"].(string),
		Name:     args["name"].(string),
		AuthorId: args["authorId"].(string),
	}
}

var Books BooksClient

func Get() BooksClient {
	return Books
}

func init() {
	Books = NewBooksClient(pg.GrpcConnection("localhost:50051"))
	queries = append(queries, &graphql.Field{
		Name: "Books_GetBooks",
		Type: GetBooksResponse_type,
		Args: GetBooksRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return Books.GetBooks(p.Context, GetBooksRequest_from_args(p.Args))
		},
	})

}
