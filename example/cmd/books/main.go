package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kitt-technology/protoc-gen-graphql/example/books"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	books.RegisterBooksServer(s, &BookService{})
	reflection.Register(s)

	fmt.Println("Serving book service on localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type BookService struct {
	books.UnimplementedBooksServer
}

func (s BookService) DoNothing(ctx context.Context, nothing *books.DoNothing) (*books.DoNothing, error) {
	return &books.DoNothing{}, nil
}

func (s BookService) GetBooksByAuthor(ctx context.Context, request *graphql.BatchRequest) (*books.GetBooksByAuthorResponse, error) {
	var bs = make(map[string]*books.BooksByAuthor)

	for _, authorId := range request.Keys {
		booksByAuthor := books.BooksByAuthor{}
		for _, book := range booksDb {
			if book.AuthorId == authorId {
				booksByAuthor.Results = append(booksByAuthor.Results, book)
			}
		}

		bs[authorId] = &booksByAuthor
	}

	return &books.GetBooksByAuthorResponse{Results: bs}, nil
}

func (s BookService) GetBooksBatch(ctx context.Context, request *books.GetBooksBatchRequest) (*books.GetBooksBatchResponse, error) {
	var res books.GetBooksBatchResponse
	for _, req := range request.Reqs {
		key := graphql.ProtoKey(req)
		resp, err := s.GetBooks(ctx, req)
		if err != nil {
			res.Results[key] = resp
		}
	}
	return &res, nil
}

func (s BookService) GetBooks(ctx context.Context, request *books.GetBooksRequest) (*books.GetBooksResponse, error) {
	var bs []*books.Book

	if request.HardbackOnly != nil && request.HardbackOnly.GetValue() {
		bs = append(bs, booksDb["3"])
	} else if len(request.Ids) > 0 {
		for _, id := range request.Ids {
			for _, b := range booksDb {
				if b.Id == id {
					bs = append(bs, b)
				}
			}
		}
	} else {
		for _, book := range booksDb {
			b := book
			bs = append(bs, b)
		}
	}

	return &books.GetBooksResponse{Books: bs}, nil
}

var booksDb map[string]*books.Book

func init() {
	var ten int64 = 10
	booksDb = map[string]*books.Book{
		"1": {
			Id:       "1",
			Name:     "Philosophers Stone",
			AuthorId: "3",
			Genre:    0,
			IsSigned: wrapperspb.Bool(true),
		},
		"2": {
			Id:       "2",
			Name:     "Chamber of Secrets ",
			AuthorId: "3",
			Genre:    1,
			IsSigned: wrapperspb.Bool(false),
		},
		"3": {
			Id:       "3",
			Name:     "Prisoner of Azkaban",
			AuthorId: "3",
			Genre:    0,
		},
		"4": {
			Id:       "4",
			Name:     "The Kreutzer Sonata",
			AuthorId: "1",
			Genre:    0,
		},
		"5": {
			Id:          "5",
			Name:        "James and the Giant Peach",
			AuthorId:    "2",
			Genre:       1,
			ReleaseDate: &timestamppb.Timestamp{Seconds: 100592},
		},
		"6": {
			Id:       "6",
			Name:     "The BFG",
			AuthorId: "2",
			Genre:    1,
			Copies:   &ten,
			Price:    10,
		},
	}
}
