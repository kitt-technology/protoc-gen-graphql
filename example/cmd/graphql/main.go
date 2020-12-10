package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/kitt-technology/protoc-gen-graphql/example/authors"
	"github.com/kitt-technology/protoc-gen-graphql/example/books"
	"log"
	"net/http"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	fields := append(authors.Fields, books.Fields...)
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
		result := graphql.Do(graphql.Params{
			Context:        req.Context(),
			Schema:         schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		})
		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	fmt.Println("Serving graphql")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func init() {
	books.Book_type.AddFieldConfig("author", &graphql.Field{
		Name: "author",
		Type: authors.Author_type,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return authors.LoadAuthor(p.Context, p.Source.(*books.Book).AuthorId)
		},
	})
	authors.Author_type.AddFieldConfig("books", &graphql.Field{
		Name: "books",
		Type: graphql.NewList(books.Book_type),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return books.LoadBook(p.Context, p.Source.(*authors.Author).Id)
		},
	})
}
