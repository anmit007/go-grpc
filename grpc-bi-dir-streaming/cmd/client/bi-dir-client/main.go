package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/anmit007/go-grpc/grpc-streaming/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewStreamingServiceClient(conn)
	stream, err := client.Echo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	eg, ctx := errgroup.WithContext(ctx)
	// create sep go routine to liisten to the server reponses
	eg.Go(func() error {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			log.Printf("Message recieved client: %v", res.GetMessage())
		}
		return nil
	})

	// send msg from client

	for i := 0; i < 5; i++ {
		req := &proto.EchoRequest{
			Message: fmt.Sprintf("Hello %d", i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 2)

	}
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Println("bi-dir-stream closed")

}
