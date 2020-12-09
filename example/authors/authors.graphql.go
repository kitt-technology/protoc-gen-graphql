package authors

import "github.com/graphql-go/graphql"
import pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
import "google.golang.org/protobuf/proto"
import "context"

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

var GetAuthorsRequest_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsRequest",
	Fields: graphql.Fields{
		"ids": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
		},
	},
})

var GetAuthorsRequest_args = graphql.FieldConfigArgument{
	"ids": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
	},
}

func GetAuthorsRequest_from_args(args map[string]interface{}) *GetAuthorsRequest {
	return &GetAuthorsRequest{
		Ids: args["ids"].([]string),
	}
}

var GetAuthorsResponse_type = graphql.NewObject(graphql.ObjectConfig{
	Name: "GetAuthorsResponse",
	Fields: graphql.Fields{
		"authors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(Author_type)),
		},
	},
})

var GetAuthorsResponse_args = graphql.FieldConfigArgument{
	"authors": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.NewList(Author_type)),
	},
}

func GetAuthorsResponse_from_args(args map[string]interface{}) *GetAuthorsResponse {
	return &GetAuthorsResponse{
		Authors: args["authors"].([]*Author),
	}
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
	return &Author{
		Id:   args["id"].(string),
		Name: args["name"].(string),
	}
}

var Authors AuthorsClient

func init() {
	Authors = NewAuthorsClient(pg.GrpcConnection("localhost:50052"))
	queries = append(queries, &graphql.Field{
		Name: "Authors_GetAuthors",
		Type: GetAuthorsResponse_type,
		Args: GetAuthorsRequest_args,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return Authors.GetAuthors(p.Context, GetAuthorsRequest_from_args(p.Args))
		},
	})

	if dataloadersToProvide == nil {
		dataloadersToProvide = make(map[string]pg.Dataloader)
	}
	dataloadersToProvide["author_id_loader"] = pg.Dataloader{
		Fn: func(ctx context.Context, ids []string) (interface{}, error) {
			resp, err := Authors.GetAuthors(ctx, &GetAuthorsRequest{Ids: ids})

			if err != nil {
				return nil, err
			}
			return resp.Authors, nil
		},
		Output: graphql.NewList(Author_type),
	}

}
