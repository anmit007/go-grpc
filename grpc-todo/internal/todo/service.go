package todo

import (
	"context"

	"github.com/anmit007/go-grpc-todo/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedTodoServiceServer
	tasks map[string]string
}

func NewService() *Service {
	return &Service{
		tasks: make(map[string]string),
	}
}

func (s *Service) AddTask(ctx context.Context, request *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {

	if request.GetTask() == "" {
		return nil, status.Error(codes.InvalidArgument, "Task cannot be empty")
	}

	id := uuid.New().String()
	s.tasks[id] = request.GetTask()

	return &proto.AddTaskResponse{
		Id: id,
	}, nil

}

func (s *Service) CompleteTask(ctx context.Context, request *proto.CompleteTaskRequest) (*proto.CompleteTaskResponse, error) {
	_, ok := s.tasks[request.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "task not found")
	}
	delete(s.tasks, request.GetId())
	return &proto.CompleteTaskResponse{}, nil
}

func (s *Service) ListTasks(ctx context.Context, request *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	tasks := make([]*proto.Task, 0, len(s.tasks))
	for id, task := range s.tasks {
		tasks = append(tasks, &proto.Task{
			Id:   id,
			Task: task,
		})
	}
	return &proto.ListTasksResponse{
		Tasks: tasks,
	}, nil
}
