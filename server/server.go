package main

import (
	"context"
	"fmt"
	"log"

	"greeting-service/api"

	"google.golang.org/grpc/metadata"
)

type server struct {
	api.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &api.HelloResponse{Message: message}, nil
}
