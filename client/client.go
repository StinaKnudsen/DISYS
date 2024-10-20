package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//ChittyChatty "mandact3.go/grpc"
	"mandact3.go/grpc/mandact3.go/grpc/ChittyChatty"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Nor working")
	}

	client := ChittyChatty.NewChittyChatClient(conn)

	messages, err := client.GetMessages(context.Background(), &ChittyChatty.Empty{})
	if err != nil {
		log.Fatalf("Nor working")
	}

	for _, message := range messages.Messages {
		println(message)
	}
}


