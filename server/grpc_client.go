// package server

// import (
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	pb "go-fiber-grpc/proto/books"
// )

// func StartGrpcREST() {
// 	// Use transport credentials instead of deprecated WithInsecure
// 	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to gRPC: %v", err)
// 	}
// 	defer conn.Close()

// 	grpcClient := pb.NewBookServiceClient(conn)

// 	app := fiber.New()

// 	app.Get("/grpc-books", func(c *fiber.Ctx) error {
// 		resp, err := grpcClient.ListBooks(c.Context(), &pb.Empty{})
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 		}
// 		return c.JSON(resp)
// 	})

// 	log.Println("🌐 REST server started on port 3001")
// 	app.Listen(":3001")
// }

package server

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-fiber-grpc/proto/books"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGrpcREST() {
	// Create a context with a timeout
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use NewClient instead of DialContext (updated method)
	conn, err := grpc.NewClient(
		"localhost:50051", // Target address
		grpc.WithTransportCredentials(insecure.NewCredentials()), // No TLS (insecure connection)
		grpc.WithBlock(),                      // Block until connection is established
		grpc.WithTimeout(10*time.Second),      // Use the context with timeout
		grpc.WithInsecure(),                   // Deprecated, but included for backward compatibility
		grpc.WithAuthority("localhost:50051"), // Authority for the connection
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	grpcClient := pb.NewBookServiceClient(conn)

	// Set up Fiber app for REST
	app := fiber.New()

	// Define REST routes and use gRPC client to interact with gRPC services
	app.Get("/grpc-books", func(c *fiber.Ctx) error {
		// Make gRPC call to fetch books
		resp, err := grpcClient.ListBooks(c.Context(), &pb.Empty{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(resp)
	})

	fmt.Println("🌐 REST server started on port 3001")
	// Start Fiber REST server on port 3001
	app.Listen(":3001")
}
