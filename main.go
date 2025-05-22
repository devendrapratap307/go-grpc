package main

import "go-fiber-grpc/server"

func main() {
	// Start REST server
	// server.StartREST()

	// Start gRPC server
	go server.StartGRPC() // Start gRPC server in a goroutine to avoid blocking

	// gRPC client
	server.StartGrpcREST()
}
