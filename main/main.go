package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	pb "supermaple.cool/maple_mobile_server/messaging"
)

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
	var srv = &MapleServer{
		clients:    map[int32]Client{},
		eventQueue: make(chan *pb.RequestEvent),
		currentCharId: 1,
	} // register the maple service to the server
	pb.RegisterMapleServiceServer(s, srv)

	service.RegisterChannelzServiceToServer(s)

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return
	}

}

type Client struct {
	charId          int32
	networkHandler  pb.MapleService_EventsStreamServer
}


type MapleServer struct {
	currentCharId int32
	clients       map[int32]Client
	mu            sync.RWMutex
	eventQueue    chan *pb.RequestEvent
}

func (s * MapleServer) sendToClient (charId int32, resp *pb.ResponseEvent) {
	log.Printf("%s - broadcasted %+v", resp, charId)
	//fixme: there is null pointer bag here
	s.clients[charId].networkHandler.Send(resp)
}

func (s *MapleServer) broadcast(resp *pb.ResponseEvent, doNotSendTO int32) {
	s.mu.Lock()
	for id, client := range s.clients {
		if id == doNotSendTO {
			continue
		}
		if client.networkHandler == nil {
			continue
		}
		if err := client.networkHandler.Send(resp); err != nil {
			log.Printf("%s - broadcast error %v", id, err)
			//currentClient.done <- errors.New("failed to broadcast message")
			continue
		}
		log.Printf("%s - broadcasted %+v", resp, id)
	}
	s.mu.Unlock()
}

func (s *MapleServer) handlePressButton(button *pb.PressButton, charId int32) {
	button.Charid = charId
	resp := pb.ResponseEvent{
		Event: &pb.ResponseEvent_PressButton{
			PressButton: button,
		},
	}
	s.broadcast(&resp, 0)
}

func (s *MapleServer) handleExpressionButton(button *pb.ExpressionButton, charId int32) {
	button.Charid = charId
	resp := pb.ResponseEvent{
		Event: &pb.ResponseEvent_ExpressionButton{
			ExpressionButton: button,
		},
	}
	s.broadcast(&resp, 0)
}

func (s *MapleServer) handlePlayerConnect(playerConnection *pb.RequestPlayerConnect, charId int32) {
	senderResp := pb.ResponseEvent{
		Event: &pb.ResponseEvent_PlayerConnected {
			PlayerConnected: &pb.ResponsePlayerConnected{
				Charid: charId,
				Hair:   30066,
				Skin:   0,
				Face:   20000,
			},
		},
	}

	broadcastResp := pb.ResponseEvent{
		Event: &pb.ResponseEvent_OtherPlayerConnected {
			OtherPlayerConnected: &pb.ResponseOtherPlayerConnected{
				Charid: charId,
				Hair:   30066,
				Skin:   0,
				Face:   20000,
			},
		},
	}


	s.sendToClient(charId, &senderResp)
	s.broadcast(&broadcastResp, charId)
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
	s.broadcast(&resp, 0)
}

func (s *MapleServer) EventsStream(server pb.MapleService_EventsStreamServer) error {
	client := Client{
		charId:         s.currentCharId,
		networkHandler: server,
	}

	s.clients[s.currentCharId] = client
	s.currentCharId = rand.Int31() //todo: change

	log.Printf("new client\n clients: %v\n", s.clients)
	go func() {
		for {
			log.Printf("wating for event\n")
			event := <-s.eventQueue
			log.Printf("got event from queue: %v", event)

			switch event.GetEvent().(type) {
				case *pb.RequestEvent_DropItem:
					s.handleDropItem(event.GetDropItem())
				case *pb.RequestEvent_PressButton:
					s.handlePressButton(event.GetPressButton(), client.charId)
				case *pb.RequestEvent_ExpressionButton:
					s.handleExpressionButton(event.GetExpressionButton(), client.charId)
				case *pb.RequestEvent_PlayerConnect:
					s.handlePlayerConnect(event.GetPlayerConnect(), client.charId)

			}
		}
	}()

	for {
		req, err := server.Recv()
		if err != nil {
			log.Printf("receive error %v", err)
			delete(s.clients, client.charId)
			return nil
			//currentClient.done <- errors.New("failed to receive request")
			//return
		}
		log.Printf("got message %+v", req)
		s.eventQueue <- req
		log.Printf("pushed to queue")
	}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mainServer := MainServer{
		grpcServerAddr: "0.0.0.0:80",
	}

	startMainServer(mainServer)
}
