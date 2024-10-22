package main

import (
	proto "ChittyServer/grpc"
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedChittyChattyServiceServer
	participants map[string]proto.ChittyChattyService_ListenToMessagesServer
	lamportClock int64
	mutex        sync.Mutex
}

// Our main <33
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterChittyChattyServiceServer(grpcServer, &server{
		participants: make(map[string]proto.ChittyChattyService_ListenToMessagesServer),
	})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Join(context context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {

	s.increment()

	for id, stream := range s.participants {
		if id != req.ParticipantId {
			broadcastMessage := &proto.BroadcastMessageRequest{
				ParticipantId:    req.ParticipantId,
				Message:          "joined Chitty-Chat at Lamport time: ",
				LogicalTimestamp: s.lamportClock,
			}
			if err := stream.Send(broadcastMessage); err != nil {
				log.Printf("failed to broadcast welcome message to %s", id)
			}
		}
	}

	return &proto.JoinResponse{
		WelcomeMessage:   "Welcome to Chitty-Chat, you joind at: ",
		LogicalTimestamp: s.lamportClock,
	}, nil
}

func (s *server) Leave(context context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {

	s.increment()
	delete(s.participants, req.ParticipantId)

	for id, stream := range s.participants {
		if id != req.ParticipantId {
			broadcastMessage := &proto.BroadcastMessageRequest{
				ParticipantId:    req.ParticipantId,
				Message:          "left Chitty-Chat at Lamport time: ",
				LogicalTimestamp: s.lamportClock,
			}
			if err := stream.Send(broadcastMessage); err != nil {
				log.Printf("failed to broadcast goodbye message to %s", id)
			}
		}
	}

	return &proto.LeaveResponse{
		GoodbyeMessage:   " left Chitty-Chat at: ",
		LogicalTimestamp: s.lamportClock,
	}, nil
}

func (s *server) PublishMessage(context context.Context, req *proto.ChatMessageRequest) (*proto.PublishResponse, error) {

	s.increment()

	for id, stream := range s.participants {
		if id != req.ParticipantId {
			broadcastMessage := &proto.BroadcastMessageRequest{
				ParticipantId:    req.ParticipantId,
				Message:          req.Message,
				LogicalTimestamp: s.lamportClock,
			}
			if err := stream.Send(broadcastMessage); err != nil {
				log.Printf("failed to publish message to %s", id)
			}
		}
	}

	return &proto.PublishResponse{Success: true}, nil
}

func (s *server) ListenToMessages(req *proto.ListenRequest, stream proto.ChittyChattyService_ListenToMessagesServer) error {

	s.participants[req.ParticipantId] = stream

	<-stream.Context().Done()

	delete(s.participants, req.ParticipantId)
	return nil
}

// Functions relating to Lamport clock
func (s *server) increment() {
	s.mutex.Lock()
	s.lamportClock++
	s.mutex.Unlock()
}
