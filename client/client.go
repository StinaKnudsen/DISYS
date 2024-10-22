package main

import (
	proto "ChittyServer/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewChittyChattyServiceClient(conn)

	joinResp, err := client.Join(context.Background(), &proto.JoinRequest{ParticipantId: "client1"})
	if err != nil {
		log.Fatalf("Join failed: %v", err)
	}
	log.Printf("Join response: %s", joinResp.WelcomeMessage)
}
