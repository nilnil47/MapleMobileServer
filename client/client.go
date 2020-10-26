package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"

	//"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	pb "supermaple.cool/maple_mobile_server/messaging"
)

var serverConnectionString = "137.135.90.47:80"

// var serverConnectionString = "localhost:80"

var grpcClient pb.MapleServiceClient

func clientSetup() pb.MapleServiceClient {
	requestOpts := grpc.WithInsecure()
	// Dial the server, returns a client connection
	conn, err := grpc.Dial(serverConnectionString, requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	client := pb.NewMapleServiceClient(conn)
	return client
}

func setup() {
	// start the server on a different thread which wont stuck the testing thread
	//go startMainServer(mainServer)
}

func main() {
	//var name string
	client := clientSetup()
	stream, err := client.EventsStream(context.TODO())
	if err != nil {
		log.Fatalf("stream error")
	}
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				fmt.Printf("error in recive: %v\n", err)
				return
			}
			fmt.Printf("recive: %v\n	", res)
		}
	}()
	for {
		scanner.Scan()

		_ = scanner.Text()
		d := pb.RequestDropItem{
			Id:    int32(rand.Uint32()),
			Count: 13,
			Owner: 4,
			Start: &pb.Point{
				X: 5,
				Y: 6,
			},
			Invtype: 11,
			Slotid:  12,
		}
		log.Printf("sends request %v", &d)
		err = stream.Send(&pb.RequestEvent{Event: &pb.RequestEvent_DropItem{DropItem: &d}})
		if err != nil {
			log.Print(err)
			return
		}
	}
}
