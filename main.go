package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
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
	srv := &MapleServer{}

	// register the maple service to the server
	pb.RegisterMapleServiceServer(s, srv)

	service.RegisterChannelzServiceToServer(s)

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

type MapleServer struct{
	clients []int
	stream pb.MapleService_EventsStreamServer
}

func (s MapleServer) broadcast(event * pb.ResponseEvent) {

	s.stream.Send(event)
}

func (s MapleServer) Connect(ctx context.Context, request *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	println("connect")
	return &pb.ConnectResponse{}, nil
}

func (MapleServer) EventsStream(server pb.MapleService_EventsStreamServer) error {
	requestEvent, err := server.Recv()
	//eventque.push(requestEvent)
	if err != nil {
		log.Fatalf("error in event stream:%v", err)
	}

	d := pb.DropItem{
		Id: requestEvent.GetDropItem().Id,
		X:  requestEvent.GetDropItem().X,
		Y:  requestEvent.GetDropItem().Y,
	}

	fmt.Print(d)
	_ = server.Send(&pb.ResponseEvent{Event: &pb.ResponseEvent_DropItem{DropItem: &d}})

	//dropItem := requestEvent.GetDropItem()
	//server.Send(pb.ResponseEvent{
	//	Event: pb.DropItem{
	//	Id: 0,
	//	X:  0,
	//	Y:  0,
	//}})

	return err
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
