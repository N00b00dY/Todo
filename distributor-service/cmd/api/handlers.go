package main

import (
	"context"
	"distributor/dbs"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Todo struct {
	ID        int       `json:"id"`
	Todo      string    `json:"todo"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// addToDo sends the toDo entry to the DB service
func (app *Config) AddToDo(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Todo   string `json:"todo"`
		Active int32  `json:"active"`
	}

	// get the payload from the request into requestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	//give the todo to the DB service to save it
	conn, err := grpc.Dial(fmt.Sprintf("%s:5001", app.dbServiceName), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	defer conn.Close()

	c := dbs.NewDbServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.AddToDo(ctx, &dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     1,
			Todo:   requestPayload.Todo,
			Active: requestPayload.Active,
		},
	})

	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}

// DeleteToDo sends ID to the DB service
func (app *Config) DeleteToDo(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int32 `json:"ID"`
	}

	// get the payload from the request into requestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:5001", app.dbServiceName), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	defer conn.Close()

	c := dbs.NewDbServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.DeleteToDo(ctx, &dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     int32(requestPayload.ID),
			Todo:   "",
			Active: 0,
		},
	})

	if err != nil {
		fmt.Println(err)
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	app.writeJSON(w, http.StatusAccepted, payload)

}

// CheckToDo sends ID to the DB service
func (app *Config) CheckToDo(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int32 `json:"ID"`
	}

	// get the payload from the request into requestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:5001", app.dbServiceName), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	defer conn.Close()

	c := dbs.NewDbServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.CheckToDo(ctx, &dbs.TodoRequest{
		TodoEntry: &dbs.Todo{
			ID:     int32(requestPayload.ID),
			Todo:   "",
			Active: 0,
		},
	})

	if err != nil {
		fmt.Println(err)
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)

}
