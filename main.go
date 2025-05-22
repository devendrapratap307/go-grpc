package main

import "go-fiber-grpc/server"

func main() {
	// Start REST server
	// server.StartREST()

	// Start gRPC server
	server.StartGRPC()

	// gRPC client
	// server.StartGrpcREST()
}
