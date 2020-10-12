package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "supermaple.cool/maple_mobile_server/messaging"
)

func main() {
	requestOpts := grpc.WithInsecure()
	// Dial the server, returns a client connection
	conn, err := grpc.Dial("192.168.1.18:50051", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	client := pb.NewMapleServiceClient(conn)
	stream, err := client.StartChatMessageStreaming(context.TODO(), &pb.Empty{})

	//for {
	// stream.Recv returns a pointer to a ListBlogRes at the current iteration
	res, err := stream.Recv()
	// If end of stream, break the loop
	//if err == io.EOF {
	//	break
	//}
	// if err, return an error
	if err != nil {
		log.Fatal("stream error")
	}
	// If everything went well use the generated getter to print the blog message
	fmt.Println(res.Message)
	//}
}
