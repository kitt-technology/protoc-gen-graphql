package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Create modules with dial options - connections will be created lazily
	productsModule := products.NewProductsModule(
		products.WithDialOptions(opts...),
	)
	usersModule := users.NewUsersModule(
		users.WithDialOptions(opts...),
	)

	// Setup cross-service relationships (e.g., add seller field to products)
	productsModule.AddFieldToProduct("seller", &graphql.Field{
		Type: usersModule.UserType(),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := p.Source.(*products.Product)
			thunk, err := usersModule.UsersLoadUsers(p, product.SellerId)
			if err != nil {
				return nil, err
			}
			return thunk()
		},
	})

	// Combine fields from all modules using helper function
	field := pg.CombineModuleFields(productsModule, usersModule)

	// Initialize context with dataloaders from all modules
	ctx := pg.WithAllLoaders(context.Background(), productsModule, usersModule)

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
	fmt.Println("========================================")
	fmt.Printf("GraphQL Server running on http://localhost:%s/graphql\n", port)
	fmt.Println("========================================")
	fmt.Println("\nTest the API with curl:")
	fmt.Println("\n1. Get all users (demonstrates pagination):")
	fmt.Println(`  curl -X POST http://localhost:8080/graphql \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"query": "{ users_GetUsers { users { id email firstName lastName } pageInfo { totalCount } } }"}'`)
	fmt.Println("\n2. Get all products (demonstrates categories, pricing, inventory):")
	fmt.Println(`  curl -X POST http://localhost:8080/graphql \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"query": "{ products_GetProducts { products { id name category price { units currencyCode } inventory { quantity } } } }"}'`)
	fmt.Println("\n3. Search products (demonstrates filtering):")
	fmt.Println(`  curl -X POST http://localhost:8080/graphql \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"query": "{ products_SearchProducts(query: \"phone\", limit: 5) { products { id name price { units } } pageInfo { hasNextPage } } }"}'`)
	fmt.Println("\n4. Get user profile (demonstrates nested data):")
	fmt.Println(`  curl -X POST http://localhost:8080/graphql \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"query": "{ users_GetUserProfile(userId: \"1\") { userId addresses { city stateProvince } loyalty { tier points } } }"}'`)
	fmt.Println("\n5. Get products with seller info (demonstrates cross-service relationships):")
	fmt.Println(`  curl -X POST http://localhost:8080/graphql \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"query": "{ products_GetProducts { products { id name seller { id email firstName lastName } } } }"}'`)
	fmt.Println("========================================")

	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}
