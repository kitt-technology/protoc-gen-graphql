package main

import (
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/example/authors"
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

func (a AuthorService) GetAuthors(ctx context.Context, request *authors.GetAuthorsRequest) (*authors.GetAuthorsResponse, error) {
	var as []*authors.Author

	for _, id := range request.Ids {
		if author, ok := authorsDb[id]; ok {
			as = append(as, &author)
		}
	}
	return &authors.GetAuthorsResponse{Authors: as}, nil
}

var authorsDb map[string]authors.Author

func init() {
	authorsDb = map[string]authors.Author{
		"1": {
			Id:   "1",
			Name: "authorsDb",
		},
		"2": {
			Id:   "3",
			Name: "Road Dahl",
		},
		"3": {
			Id:   "3",
			Name: "J.K. Rowling",
		},
	}
}
