package main

import (
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/example/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

func (s BookService) GetBooksByAuthor(ctx context.Context, request *books.GetBooksRequest) (*books.GetBooksByAuthorResponse, error) {
	var bs []*books.BooksByAuthor

	for _, authorId := range request.Ids {
		booksByAuthor := books.BooksByAuthor{AuthorId: authorId}
		for _, book := range booksDb {
			if book.AuthorId == authorId {
				booksByAuthor.Books = append(booksByAuthor.Books, book)
			}
		}

		bs = append(bs, &booksByAuthor)
	}

	return &books.GetBooksByAuthorResponse{Results: bs}, nil
}

func (s BookService) GetBooks(ctx context.Context, request *books.GetBooksRequest) (*books.GetBooksResponse, error) {
	var bs []*books.Book

	if len(request.Ids) > 0 {
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
	booksDb = map[string]*books.Book{
		"1": {
			Id:       "1",
			Name:     "Philosophers Stone",
			AuthorId: "3",
		},
		"2": {
			Id:       "2",
			Name:     "Chamber of Secrets ",
			AuthorId: "3",
		},
		"3": {
			Id:       "3",
			Name:     "Prisoner of Azkaban",
			AuthorId: "3",
		},
		"4": {
			Id:       "4",
			Name:     "The Kreutzer Sonata",
			AuthorId: "1",
		},
		"5": {
			Id:       "5",
			Name:     "James and the Giant Peach",
			AuthorId: "2",
		},
		"6": {
			Id:       "6",
			Name:     "The BFG",
			AuthorId: "2",
		},
	}
}
