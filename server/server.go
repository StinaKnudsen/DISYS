package main

import (
	proto "ChittyServer/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedChittyChattyServiceServer
	// Unfinished
}

// Make the grpc methods

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterChittyChattyServiceServer(grpcServer, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
