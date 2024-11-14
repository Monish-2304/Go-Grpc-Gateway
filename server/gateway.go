package main

import (
	"context"
	"greeting-service/api"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	go startGRPCServer()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterGreetingServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGreetingServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
