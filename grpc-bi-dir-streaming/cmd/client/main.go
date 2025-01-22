package main

import (
	"context"
	"log"

	"github.com/anmit007/go-grpc/grpc-streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()
	// init grpc conn
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	// create the client
	client := proto.NewStreamingServiceClient(conn)

	//init stream

	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	// loop through all the responses we get back from the server
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("recieved time from server: %s", res.CurrentTime.AsTime())
	}

	// once the server closes the stream, exit gracefully
	log.Println("SERVER STREAM CLOSED")
}
