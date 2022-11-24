package main

import (
	"fmt"
	"log"
	"net/http"
)

var portNumber = "8081"

type Config struct {
}

func main() {
	app := Config{}

	fmt.Printf("Server startet at port %v", portNumber)
	serv := &http.Server{
		Addr:    fmt.Sprintf(":%s", portNumber),
		Handler: app.routes(),
	}

	err := serv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
