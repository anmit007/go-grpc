package main

import (
	"anmit007/go-grpc/proto"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)
	res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Anmit"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("resposne recieved: %s", res.Message)
}
