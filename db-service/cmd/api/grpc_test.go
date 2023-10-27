package main

import (
	"context"
	"db-service/dbs"
	"testing"
)

func Test_DBServer_AddToDo(t *testing.T) {
	// create a request to pass to our handler
	todoReq := dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     1,
			Todo:   "test",
			Active: 0,
		},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "test")

	restResp, err := testDBServer.AddToDo(ctx, &todoReq)
	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}
	if restResp.Result != "added" {
		t.Errorf("Expected result to be 'added', got %s", restResp.Result)
	}
}

func Test_DBServer_CheckTodo(t *testing.T) {
	// create a request to pass to our handler
	todoReq := dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     1,
			Todo:   "test",
			Active: 0,
		},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "test")

	restResp, err := testDBServer.CheckToDo(ctx, &todoReq)

	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}
	if restResp.Result != "checked" {
		t.Errorf("Expected result to be 'added', got %s", restResp.Result)
	}
}

func Test_DBServer_GetAllTodos(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "test")

	restResp, err := testDBServer.GetAllToDos(ctx, &dbs.TodoListRequest{})

	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}
	if restResp.Result != "Got all todos" {
		t.Errorf("Expected result to be 'added', got %s", restResp.Result)
	}
}

func Test_DBServer_DeleteToDo(t *testing.T) {
	// create a request to pass to our handler
	todoReq := dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     1,
			Todo:   "test",
			Active: 0,
		},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "test")

	restResp, err := testDBServer.DeleteToDo(ctx, &todoReq)
	if err != nil {
		t.Errorf("Error adding todo: %v", err)
	}
	if restResp.Result != "deleted" {
		t.Errorf("Expected result to be 'added', got %s", restResp.Result)
	}
}
