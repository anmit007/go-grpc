package main

import (
	"anmit007/go-grpc/proto"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: ""})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code %s, error:%s", status.Code().String(), status.Message())
		}

		log.Fatal(err)
	}
	log.Printf("resposne recieved: %s", res.Message)
}
