package main

import (
	"database/sql"
	"db-service/data"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "5001"
)

var counts int64

type Config struct {
	Repo   *data.PostgresRepository
	Client *http.Client
}

func main() {

	// connect to DB
	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	app := Config{
		Client: &http.Client{},
		Repo:   data.NewPostgresRepository(conn),
	}

	go app.gRPCListen()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	fmt.Println(dsn)
	for {
		connection, err := openDB(dsn)
		if err != nil {
			fmt.Println(err)
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
