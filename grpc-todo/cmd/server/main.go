package main

import (
	"log"
	"net"

	"github.com/anmit007/go-grpc-todo/internal/todo"
	"github.com/anmit007/go-grpc-todo/proto"
	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	todoService := todo.NewService()
	proto.RegisterTodoServiceServer(grpcServer, todoService)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("STARTING SERVER ON ADDRESS %s ", ":50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
