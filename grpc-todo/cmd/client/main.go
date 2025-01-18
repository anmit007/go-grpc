package main

import (
	"context"
	"log"

	"github.com/anmit007/go-grpc-todo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)
	task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "Linear Algerba"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Added task -id:%s", task1.GetId())
	task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "ML spec"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Added task -id:%s", task2.GetId())

	allTasks, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listing all tasks: %v", allTasks.GetTasks())

	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{
		Id: task1.GetId(),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("completed task - id: %s", task1.GetId())

}
