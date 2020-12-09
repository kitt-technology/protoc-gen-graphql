package main

import (
	"encoding/json"
	"fmt"
	"github.com/kitt-technology/protoc-gen-graphql/example/authors"
	"github.com/kitt-technology/protoc-gen-graphql/example/books"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func AddFieldResolver(fieldName string, root *graphql.Object, returnType graphql.Output, resolve func(p graphql.ResolveParams) (interface{}, error)) {
	root.AddFieldConfig(fieldName, &graphql.Field{
		Name:    fieldName,
		Type:    returnType,
		Resolve: resolve,
	})
}

func main() {
	config := pg.ProtoConfig{}

	AddFieldResolver("author", books.Book_type, authors.Author_type, func(p graphql.ResolveParams) (interface{}, error) {
		book := p.Source.(*books.Book)
		resp, err := authors.Get().GetAuthors(p.Context, &authors.GetAuthorsRequest{
			Ids: []string{book.AuthorId},
		})

		if err != nil {
			return nil, err
		}
		return resp.Authors[0], err
	})

	config = authors.Register(config)
	config = books.Register(config)

	field := graphql.Fields{}
	for _, query := range config.Queries {
		field[query.Name] = query
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
