package main

import (
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/example/authors"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authors.RegisterAuthorsServer(s, &AuthorService{})
	reflection.Register(s)

	fmt.Println("Serving author service on localhost:50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type AuthorService struct {
	authors.UnimplementedAuthorsServer
}

func (a AuthorService) LoadAuthors(ctx context.Context, request *graphql.BatchRequest) (*authors.AuthorsBatchResponse, error) {
	as := make(map[string]*authors.Author)

	for _, key := range request.Keys {
		for _, a := range authorsDb {
			if a.Id == key {
				as[key] = a
			}
		}
	}

	return &authors.AuthorsBatchResponse{Results: as}, nil
}

func (a AuthorService) GetAuthors(ctx context.Context, request *authors.GetAuthorsRequest) (*authors.GetAuthorsResponse, error) {
	var as []*authors.Author

	if len(request.Ids) > 0 {
		for _, id := range request.Ids {
			for _, a := range authorsDb {
				if a.Id == id {
					as = append(as, a)
				}
			}
		}
	} else {
		for _, author := range authorsDb {
			a := author
			as = append(as, a)
		}
	}

	return &authors.GetAuthorsResponse{Authors: as}, nil
}

var authorsDb map[string]*authors.Author

func init() {
	authorsDb = map[string]*authors.Author{
		"1": {
			Id:   "1",
			Name: "Leo Tolstoy",
		},
		"2": {
			Id:   "2",
			Name: "Road Dahl",
		},
		"3": {
			Id:   "3",
			Name: "J.K. Rowling",
		},
	}
}
