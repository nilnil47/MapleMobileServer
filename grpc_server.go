package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "supermaple.cool/maple_mobile_server/messaging"
	"sync"
)

var once sync.Once

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	once.Do(func() { // <-- atomic, does not allow repeating

		instance = make(singleton) // <-- thread safe

	})

	return instance
}

type MapleServiceServer struct{}

func (MapleServiceServer) SendChatMessage(ctx context.Context, message *pb.ChatMessage) (*pb.Empty, error) {
	panic("implement me")
}

func (MapleServiceServer) StartChatMessageStreaming(empty *pb.Empty, stream pb.MapleService_StartChatMessageStreamingServer) error {
	for {
		fmt.Println("sending")
		stream.Send(&pb.ChatMessage{Id: 1, Message: "ss"})
	}
}

//func (s *MapleServiceServer) SendChatMessage()  {
//
//}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port :50051...")

	// 50051 is the default port for gRPC
	// Ideally we'd use 0.0.0.0 instead of localhost as well
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	// slice of gRPC options
	// Here we can configure things like TLS
	opts := []grpc.ServerOption{}
	// var s *grpc.Server
	s := grpc.NewServer(opts...)
	// var srv *BlogServiceServer
	srv := &MapleServiceServer{}

	pb.RegisterMapleServiceServer(s, srv)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
