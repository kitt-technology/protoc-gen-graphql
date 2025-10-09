package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"github.com/kitt-technology/protoc-gen-graphql/example/users"
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

	// Get service addresses from environment or use defaults
	productsAddr := os.Getenv("PRODUCTS_SERVICE_ADDR")
	if productsAddr == "" {
		productsAddr = "localhost:50051"
	}

	usersAddr := os.Getenv("USERS_SERVICE_ADDR")
	if usersAddr == "" {
		usersAddr = "localhost:50052"
	}

	// Create gRPC connections
	productsConn, err := grpc.Dial(productsAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to products service at %s: %v", productsAddr, err)
	}
	defer productsConn.Close()

	usersConn, err := grpc.Dial(usersAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to users service at %s: %v", usersAddr, err)
	}
	defer usersConn.Close()

	// Initialize services and get fields
	productsClient := products.NewProductsClient(productsConn)
	usersClient := users.NewUsersClient(usersConn)

	ctx, usersFields := users.Init(context.Background(), users.WithClient(usersClient), users.WithDialOptions(opts...))
	ctx, productsFields := products.Init(ctx, products.WithClient(productsClient), products.WithDialOptions(opts...))

	fields := append(usersFields, productsFields...)
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
	fmt.Println("========================================")
	fmt.Printf("GraphQL Server running on http://localhost:%s/graphql\n", port)
	fmt.Println("========================================")
	fmt.Println("\nConnected to:")
	fmt.Printf("  - Products service: %s\n", productsAddr)
	fmt.Printf("  - Users service: %s\n", usersAddr)
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
	fmt.Println("========================================")

	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}

func init() {
	// Cross-service relationships: Add seller info to products
	products.ProductGraphqlType.AddFieldConfig("seller", &graphql.Field{
		Type: users.UserGraphqlType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := p.Source.(*products.Product)
			return users.LoadUsers(p, product.SellerId)
		},
	})
}
