package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/kitt-technology/protoc-gen-graphql/example/authors"
	"github.com/kitt-technology/protoc-gen-graphql/example/books"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Initialize services and get fields
	ctx, authorsFields := authors.Init(context.Background(), authors.WithDialOptions(opts...))
	ctx, booksFields := books.Init(ctx, books.WithDialOptions(opts...))

	fields := append(authorsFields, booksFields...)
	field := graphql.Fields{}
	for _, f := range fields {
		field[f.Name] = f
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: field}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		var p postData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		// Initialize services and get context with dataloaders
		result := graphql.Do(graphql.Params{
			Context:        ctx,
			Schema:         schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})
		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	port := "8080"
	fmt.Printf("Serving graphql on localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}

func init() {
	books.BookGraphqlType.AddFieldConfig("author", &graphql.Field{
		Type: authors.AuthorGraphqlType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return authors.LoadAuthors(p, p.Source.(*books.Book).AuthorId)
		},
	})
	books.BookGraphqlType.AddFieldConfig("foo", &graphql.Field{
		Type: graphql.Boolean,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			b := wrapperspb.Bool(true)
			return b.Value, nil
		},
	})

	authors.AuthorGraphqlType.AddFieldConfig("books", &graphql.Field{
		Type: books.BooksByAuthorGraphqlType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return books.GetBooksByAuthor(p, p.Source.(*authors.Author).Id)
		},
	})
}
