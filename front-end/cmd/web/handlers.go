package main

import (
	"context"
	"embed"
	"fmt"
	"front-end/dbs"
	"net/http"
	"text/template"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type pageData struct {
	TodList   []*dbs.Todo
	CSRFToken string
}

//go:embed templates
var templateFS embed.FS

func (app *Config) Home(w http.ResponseWriter, r *http.Request) {
	// TODO: here it would make sense to use an seperate Render function
	p := "todo.page.gohtml"
	partials := []string{
		"templates/base.layout.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", p))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// parse a fileSystem
	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get todo list from DB service
	conn, err := grpc.Dial("db-service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	defer conn.Close()

	c := dbs.NewDbServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	todoList, err := c.GetAllToDos(ctx, &dbs.TodoListRequest{})

	if err != nil {
		app.throwJSONError(w, err, http.StatusBadRequest)
		return
	}

	// create newPageData object so CSRFToken is included
	// should check for other possibilitys
	newPageData := pageData{
		TodList: todoList.TodoList,
	}

	// Execute and render templates with data
	err = tmpl.Execute(w, newPageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
