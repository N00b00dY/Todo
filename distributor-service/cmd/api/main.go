package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	dbServiceName string
}

func main() {
	app := Config{
		dbServiceName: "db-service",
	}

	log.Printf("Starting broker service on port %s \n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
