package main

import (
	proto "ChittyServer/grpc"
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewChittyChattyServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	log.Print("Please enter your participant ID: ")
	participantId, _ := reader.ReadString('\n')
	participantId = strings.TrimSpace(participantId)

	joinResp, err := client.Join(context.Background(), &proto.JoinRequest{ParticipantId: participantId})
	if err != nil {
		log.Fatalf("Join failed: %v", err)
	}
	log.Printf("Join response: %s", joinResp.WelcomeMessage)

	go ListenForMessages(client, participantId)

	for {
		log.Print("Enter a message or type 'exit' to leave :))")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			leaveChat(client, participantId)
			break
		}

		publishResp, err := client.PublishMessage(context.Background(), &proto.ChatMessageRequest{
			ParticipantId:    participantId,
			Message:          message,
			LogicalTimestamp: time.Now().Unix(), // CHANGE TO LAMPORT FOR REAL
		})
		if err != nil {
			log.Printf("failed to publish message")
		} else if publishResp.Success {
			log.Println("You have succesfully published a message ;)")
		}
	}
}

func ListenForMessages(client proto.ChittyChattyServiceClient, paticipantId string) {

	stream, err := client.ListenToMessages(context.Background(), &proto.ListenRequest{
		ParticipantId: paticipantId,
	})
	if err != nil {
		log.Fatalf("failed to listen to message")
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving message")
		}
		log.Printf("(Message from %s at Lamport time %d): %s", message.ParticipantId, message.LogicalTimestamp, message.Message)
	}
}

func leaveChat(client proto.ChittyChattyServiceClient, participantId string) {
	leaveResp, err := client.Leave(context.Background(), &proto.LeaveRequest{
		ParticipantId: participantId,
	})
	if err != nil {
		log.Fatal("failed to leave chat :()")
	}

	log.Printf("%s %s %d", participantId, leaveResp.GoodbyeMessage, leaveResp.LogicalTimestamp)

}
