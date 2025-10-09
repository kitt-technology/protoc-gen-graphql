package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/products/internal/repository"
	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/products/internal/service"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize repository
	repo := repository.NewInMemoryProductRepository()

	// Initialize service
	productService := service.NewProductService(repo)

	// Create gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	products.RegisterProductsServer(server, productService)
	reflection.Register(server)

	// Print startup information
	printStartupInfo()

	// Start serving
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func printStartupInfo() {
	fmt.Println("========================================")
	fmt.Println("Products gRPC Service")
	fmt.Println("========================================")
	fmt.Println("Listening on: localhost:50051")
	fmt.Println("\nThis service provides product catalog data")
	fmt.Println("Available sample products:")
	fmt.Println("  - Wireless Bluetooth Headphones ($79.99)")
	fmt.Println("  - Running Shoes - Trail Edition ($129.99)")
	fmt.Println("  - Smart Watch Pro ($299.99)")
	fmt.Println("  - Premium Coffee Maker ($89.99)")
	fmt.Println("  - Eco-Friendly Yoga Mat ($29.99)")
	fmt.Println("\nReady to accept gRPC requests...")
	fmt.Println("========================================")
}
