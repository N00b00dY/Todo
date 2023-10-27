package main

import (
	"context"
	"db-service/data"
	"db-service/dbs"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type DBServer struct {
	dbs.UnimplementedDbServiceServer
	PostgresRepository data.PostgresRepository // not right
}

func (l *DBServer) AddToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {
	input := req.GetTodoEntry()

	// write the log
	todoEntry := data.Todo{
		ID:        1,
		Todo:      input.Todo,
		Active:    int(input.Active),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := l.PostgresRepository.Insert(todoEntry)
	if err != nil {
		fmt.Println("Insert not working")
		res := &dbs.TodoResponse{Result: "failed"}
		return res, err
	}

	res := &dbs.TodoResponse{Result: "added"}
	return res, nil

}

func (l *DBServer) DeleteToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {
	input := req.GetTodoEntry()
	err := l.PostgresRepository.DeleteByID(int(input.ID))
	if err != nil {
		fmt.Println("Could not delete todo")
		res := &dbs.TodoResponse{Result: "failed"}
		return res, err
	}

	res := &dbs.TodoResponse{Result: "deleted"}
	return res, nil

}

func (l *DBServer) CheckToDo(ctx context.Context, req *dbs.TodoRequest) (*dbs.TodoResponse, error) {
	input := req.GetTodoEntry()
	todo, err := l.PostgresRepository.GetOne(int(input.ID))
	if err != nil {
		fmt.Println("Could not get todo")
		res := &dbs.TodoResponse{Result: "failed"}
		return res, err
	}

	if todo.Active == 1 {
		todo.Active = 0
		err = l.PostgresRepository.Update(*todo)
	} else {
		todo.Active = 1
		err = l.PostgresRepository.Update(*todo)
	}
	if err != nil {
		fmt.Println("Updating todo is not working")
		res := &dbs.TodoResponse{Result: "failed"}
		return res, err
	}

	res := &dbs.TodoResponse{Result: "checked"}
	return res, nil

}

func (l *DBServer) GetAllToDos(ctx context.Context, req *dbs.TodoListRequest) (*dbs.TodoListResponse, error) {
	todoList, err := l.PostgresRepository.GetAll()
	if err != nil {
		res := &dbs.TodoListResponse{
			TodoList: nil,
			Result:   "failed"}
		return res, err
	}

	res := &dbs.TodoListResponse{
		TodoList: todoList,
		Result:   "Got all todos",
	}
	return res, nil

}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Fail to listen for gRPC: Q%v", err)

	}

	s := grpc.NewServer()

	dbs.RegisterDbServiceServer(s, &DBServer{PostgresRepository: *app.Repo})

	log.Printf("gRPC Server started on Port %s ", gRpcPort)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Fail to listen for gRPC: Q%v", err)

	}

}
