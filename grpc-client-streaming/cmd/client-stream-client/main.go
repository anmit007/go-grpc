package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anmit007/go-grpc/grpc-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// init grpc conn
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// create our client
	client := proto.NewStreamingServiceClient(conn)

	// init the client stream
	stream, err := client.LogStream(ctx)
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// send some log messages
	for i := 0; i < 5; i++ {
		req := &proto.LogStreamRequest{
			Timestamp: timestamppb.New(time.Now()),
			Level:     proto.LogLevel_LOG_LEVEL_INFO,
			Message:   fmt.Sprintf("Hello log: %v", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		time.Sleep(time.Second)
	}

	// close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	// log the response from server
	log.Printf("Number of messages sent: %v", res.EntriesLogged)
}
