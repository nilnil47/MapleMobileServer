package main

import (
	"bufio"
	"context"
	//"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "supermaple.cool/maple_mobile_server/messaging"
)


var grpcClient pb.MapleServiceClient

func clientSetup() pb.MapleServiceClient {
	requestOpts := grpc.WithInsecure()
	// Dial the server, returns a client connection
	conn, err := grpc.Dial("localhost:50051", requestOpts)
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

	for {
		scanner.Scan()

		_ = scanner.Text()
		d := pb.DropItem{
			Id: 1,
			X:  2,
			Y:  3,
		}

		_ = stream.Send(&pb.RequestEvent{Event: &pb.RequestEvent_DropItem{DropItem: &d}})
		// stream.Recv returns a pointer to a ListBlogRes at the current iteration
		//event, err := stream.Recv()
		//if err != nil {
		//	log.Fatalf("reciving error")
		//}
		//fmt.Println(event.GetDropItem())
	}

}
