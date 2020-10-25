package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
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
	srv := &MapleServer{
		clients:    map[uuid.UUID]pb.MapleService_EventsStreamServer{},
		eventQueue: make(chan *pb.RequestEvent),
	}

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

type MapleServer struct {
	clients    map[uuid.UUID]pb.MapleService_EventsStreamServer
	mu         sync.RWMutex
	eventQueue chan *pb.RequestEvent
}

func (s MapleServer) Connect(ctx context.Context, request *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	println("connect")
	return &pb.ConnectResponse{}, nil
}

func (s *MapleServer) broadcast(resp *pb.ResponseEvent) {
	s.mu.Lock()
	for id, streamServer := range s.clients {
		if streamServer == nil {
			continue
		}
		if err := streamServer.Send(resp); err != nil {
			log.Printf("%s - broadcast error %v", id, err)
			//currentClient.done <- errors.New("failed to broadcast message")
			continue
		}
		log.Printf("%s - broadcasted %+v", resp, id)
	}
	s.mu.Unlock()
}

func (s *MapleServer) handleDropItem(item *pb.RequestDropItem) {
	resp := pb.ResponseEvent{
		Event: &pb.ResponseEvent_DropItem{
			DropItem: &pb.ResponseDropItem{
				// the oid is random uint32 number for now
				Oid:   rand.Int31(),
				Id:    item.Id,
				Owner: item.Owner,
				Start: item.Start,
				Mapid: item.Mapid,
			},
		},
	}
	log.Print("send event in broadcast: %v\n", &resp)
	s.broadcast(&resp)
}

func (s *MapleServer) EventsStream(server pb.MapleService_EventsStreamServer) error {
	//todo: change this
	clientUuid := uuid.New()
	s.clients[clientUuid] = server
	log.Printf("new client\n clients: %v\n", s.clients)
	go func() {
		for {
			log.Printf("wating for event\n")
			event := <-s.eventQueue
			log.Printf("got event from queue: %v", event)
			switch event.GetEvent().(type) {
			case *pb.RequestEvent_DropItem:
				s.handleDropItem(event.GetDropItem())
			}
		}
	}()
	//go func() {
	for {
		req, err := server.Recv()
		if err != nil {
			log.Printf("receive error %v", err)
			delete(s.clients, clientUuid)
			return nil
			//currentClient.done <- errors.New("failed to receive request")
			//return
		}
		log.Printf("got message %+v", req)
		s.eventQueue <- req
		log.Printf("pushed to queue")
	}
	//}()

	//dropItem := requestEvent.GetDropItem()
	//server.Send(pb.ResponseEvent{
	//	Event: pb.DropItem{
	//	Id: 0,
	//	X:  0,
	//	Y:  0,
	//}})
	return nil
}

//func (s *MapleServiceServer) SendChatMessage()  {
//
//}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mainServer := MainServer{
		httpServerAddr: "0.0.0.0:9000",
		httpFileServer: http.FileServer(http.Dir("http_server_files")),
		grpcServerAddr: "0.0.0.0:80",
	}

	startMainServer(mainServer)
}
