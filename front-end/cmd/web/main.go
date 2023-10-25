package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = "8081"

type Config struct {
	dbServiceName string
}

func main() {
	app := Config{
		dbServiceName: "db-service",
	}

	log.Printf("Server startet at port %v", portNumber)
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%s", portNumber),
		Handler: app.routes(),
	}

	err := serv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
