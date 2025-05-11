package server

import (
	"fmt"
	"go-fiber-grpc/database"
	"go-fiber-grpc/internal/books"
	pb "go-fiber-grpc/proto/books" // generated from books.proto

	"net"

	"google.golang.org/grpc"
)

func StartGRPC() {
	// Connect to DB
	db := database.Connect()
	repo := books.NewRepository(db)
	service := books.NewService(repo)

	// Start listening on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()

	// Register your gRPC service here
	pb.RegisterBookServiceServer(grpcServer, books.NewGRPCHandler(service))

	fmt.Println("🚀 gRPC server started on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(fmt.Sprintf("Failed to serve: %v", err))
	}
}
