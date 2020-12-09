package books

import "github.com/graphql-go/graphql"
import pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
import "google.golang.org/protobuf/proto"

var mutations []*graphql.Field
var queries []*graphql.Field

var mutationResolver func(command proto.Message, success proto.Message) (proto.Message, error)
var dataloadersToRegister map[string][]pg.RegisterDataloaderFn
var dataloadersToProvide map[string]pg.Dataloader

func AppendDataloaders(dataloaders map[string]pg.Dataloader) map[string]pg.Dataloader {
	for k, v := range dataloadersToProvide {
		dataloaders[k] = v
	}
	return dataloaders
}

func Register(config pg.ProtoConfig, mr func(command proto.Message, success proto.Message) (proto.Message, error), dataloaders map[string]pg.Dataloader) pg.ProtoConfig {
	mutationResolver = mr
	config.Mutations = append(config.Mutations, mutations...)
	config.Queries = append(config.Queries, queries...)

	// Find objects who have registered a particular dataloader and add the field resolve
	for dataloaderName, dataloader := range dataloaders {
		for _, registerFn := range dataloadersToRegister[dataloaderName] {
			registerFn(dataloader)
		}
	}
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

	if dataloadersToRegister == nil {
		dataloadersToRegister = make(map[string][]pg.RegisterDataloaderFn)
	}

	if _, ok := dataloadersToRegister["author_id_loader"]; !ok {
		dataloadersToRegister["author_id_loader"] = []pg.RegisterDataloaderFn{}
	}

	dataloadersToRegister["author_id_loader"] = append(dataloadersToRegister["author_id_loader"], func(dl pg.Dataloader) {
		Book_type.AddFieldConfig("author", &graphql.Field{
			Type: dl.Output,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				source, _ := p.Source.(*Book)
				return dl.Fn(p.Context, []string{source.AuthorId})
			},
		})
	})

}
