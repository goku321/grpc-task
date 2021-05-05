package main

import (
	"context"
	"log"
	"net"

	pb "github.com/goku321/grpc-example/task"
	"google.golang.org/grpc"
)

const (
	port = ":59690"
)

type server struct {
	pb.UnimplementedTaskServer
	tasks []string
}

func (s *server) Create(ctx context.Context, in *pb.TaskRequest) (*pb.TaskReply, error) {
	s.tasks = append(s.tasks, in.GetName())
	log.Printf("created task: %s", in.GetName())
	return &pb.TaskReply{Name: "success"}, nil
}

func (s *server) Get(ctx context.Context, in *pb.TaskRequest) (*pb.TaskReply, error) {
	if search(s.tasks, in.GetName()) {
		return &pb.TaskReply{Name: "task found"}, nil
	}
	return &pb.TaskReply{Name: "task not found"}, nil
}

func search(x []string, toFind string) bool {
	for _, t := range x {
		if t == toFind {
			return true
		}
	}
	return false
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
