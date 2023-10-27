package main

import (
	"context"
	"distributor/dbs"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"testing"
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

func (l *DBServer) AddToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {
	res := &dbs.TodoResponse{
		Result: "Added",
	}
	return res, nil

}

func (l *DBServer) DeleteToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {

	res := &dbs.TodoResponse{Result: "deleted"}
	return res, nil

}

func (l *DBServer) CheckToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {

	res := &dbs.TodoResponse{Result: "checked"}
	return res, nil

}
