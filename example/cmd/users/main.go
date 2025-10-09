package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/users/internal/repository"
	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/users/internal/service"
	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize repository
	repo := repository.NewInMemoryUserRepository()

	// Initialize service
	userService := service.NewUserService(repo)

	// Create gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	users.RegisterUsersServer(server, userService)
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
	fmt.Println("Users gRPC Service")
	fmt.Println("========================================")
	fmt.Println("Listening on: localhost:50052")
	fmt.Println("\nThis service provides user/customer data")
	fmt.Println("Available sample users:")
	fmt.Println("  - alice@example.com (Gold tier customer)")
	fmt.Println("  - bob@example.com (Platinum tier seller)")
	fmt.Println("  - charlie@example.com (Silver tier customer)")
	fmt.Println("\nReady to accept gRPC requests...")
	fmt.Println("========================================")
}
