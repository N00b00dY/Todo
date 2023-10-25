package main

import (
	"context"
	"fmt"
	"front-end/dbs"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {

	go grpcServer()

	os.Exit(m.Run())
}

func grpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "5001"))

	s := grpc.NewServer()

	dbs.RegisterDbServiceServer(s, &DBServer{})

	log.Printf("gRPC Server started on Port %s ", "5001")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Fail to listen for gRPC: Q%v", err)
	}
}

type DBServer struct {
	dbs.UnimplementedDbServiceServer
}

type Todo struct {
	ID        int       `json:"id"`
	Todo      string    `json:"todo"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l *DBServer) GetAllToDos(ctx context.Context, req *dbs.TodoListRequest) (*dbs.TodoListResponse, error) {
	todoList := []*dbs.Todo{{
		ID:     1,
		Todo:   "Ich bin der erste Test",
		Active: 0,
	},
	}
	res := &dbs.TodoListResponse{
		TodoList: todoList,
		Result:   "Got all todos",
	}
	return res, nil

}
