package main

import (
	"context"
	"log"
	"time"

	pb "github.com/goku321/grpc-example/task"
	"google.golang.org/grpc"
)

const (
	address = "localhost:59690"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTaskClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	r, err := c.Create(ctx, &pb.TaskRequest{Name: "important task"})
	if err != nil {
		log.Fatalf("failed to create task\n")
	}
	log.Printf("create task status: %s\n", r.GetName())

	// retrieve task
	r, err = c.Get(ctx, &pb.TaskRequest{Name: "important task"})
	if err != nil {
		log.Fatalf("failed to retrieve task\n")
	}
	log.Printf("retreive task status: %s\n", r.GetName())
}