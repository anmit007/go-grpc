package main

import (
	"anmit007/go-grpc/internal/hello"
	"anmit007/go-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	helloService := hello.Service{}
	proto.RegisterHelloServiceServer(grpcServer, &helloService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting server on address: %s", "50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
