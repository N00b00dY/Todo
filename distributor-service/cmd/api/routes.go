package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "Options"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           30,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	// define routes
	router.Post("/addToDo", app.AddToDo)
	router.Post("/checkToDo", app.CheckToDo)
	// should have used deleted methode
	router.Post("/delete-todo", app.DeleteToDo)

	return router
}
