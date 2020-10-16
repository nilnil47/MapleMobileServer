package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	pb "supermaple.cool/maple_mobile_server/messaging"
)

//type Server interface {
//	start()
//	stop()
//}

type MainServer struct {
	httpServerAddr string
	httpFileServer http.Handler

	grpcServerAddr string
}

func startMainServer(server MainServer) {

	// create listener for grpc server
	fmt.Printf("starting grpc server on %s", server.grpcServerAddr)
	listener, err := net.Listen("tcp", server.grpcServerAddr)

	if err != nil {
		log.Fatalf("Unable to listen on port %s: %v", server.grpcServerAddr, err)
	}

	// start http files server
	//err = http.ListenAndServe(server.httpServerAddr, server.httpFileServer)
	//if err != nil {
	//	log.Fatalf("enable to start http server in addres %s", server.httpServerAddr)
	//}
	// create grpc server
	s := grpc.NewServer()

	// create maple service which handle all the rpc
	// defined in the proto file
	srv := &MapleServiceServer{}

	// register the maple service to the server
	pb.RegisterMapleServiceServer(s, srv)

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}

}

func stopMainServer() {
}

func (MainServer) stop() {
	panic("implement me")
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

	mainServer := MainServer{
		httpServerAddr: "0.0.0.0:9000",
		httpFileServer: http.FileServer(http.Dir("http_server_files")),
		grpcServerAddr:  "0.0.0.0:" + os.Getenv("PORT"),
	}

	startMainServer(mainServer)
}
