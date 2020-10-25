package tests

// import (
// 	"context"
// 	"fmt"
// 	"google.golang.org/grpc"
// 	"log"
// 	"os"
// 	pb "supermaple.cool/maple_mobile_server/messaging"
// 	"testing"
// )

// var mainServer = MainServer{
// 	grpcServerAddr: "https://secure-tor-35210.herokuapp.com/:50051",
// }

// var grpcClient pb.MapleServiceClient

// func clientSetup() pb.MapleServiceClient {
// 	requestOpts := grpc.WithInsecure()
// 	// Dial the server, returns a client connection
// 	conn, err := grpc.Dial("localhost:50051", requestOpts)
// 	if err != nil {
// 		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
// 	}
// 	client := pb.NewMapleServiceClient(conn)
// 	return client
// }

// func setup() {
// 	// start the server on a different thread which wont stuck the testing thread
// 	go startMainServer(mainServer)
// }

// func TestMain(m *testing.M) {
// 	setup()
// 	grpcClient = clientSetup()
// 	code := m.Run()
// 	os.Exit(code)
// }

// func TestMapleServiceServer_SendChatMessage(t *testing.T) {

// 	// start the StartChatMessageStreaming service
// 	// context.TODO() means that non specific context need in this connection
// 	stream, err := grpcClient.StartChatMessageStreaming(context.TODO(), &pb.Empty{})
// 	if err != nil {
// 		t.Fatalf("stream error")
// 	}

// 	// stream.Recv returns a pointer to a ListBlogRes at the current iteration
// 	res, err := stream.Recv()
// 	if err != nil {
// 		t.Fatalf("reciving error")
// 	}
// 	fmt.Println(res.Message)
// }
