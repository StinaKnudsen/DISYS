package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"mandact3.go/grpc/mandact3.go/grpc/ChittyChatty"
)

type databaseServer struct {
	ChittyChatty.UnimplementedChittyChatServer
	messages []string
}

func (m *databaseServer) GetMessages(context context.Context, in *ChittyChatty.Empty) (*ChittyChatty.Message, error) {
	return &ChittyChatty.Message{Message: m.messages}, nil
}

func main() {
	server := &databaseServer{messages: []string{}}

	server.messages = append(server.messages, "yolo")

	server.startServer()
}

func (m *databaseServer) startServer() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Did not work")
	}

	ChittyChatty.RegisterChittyChatServer(grpcServer, m)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Did not work")
	}
}
