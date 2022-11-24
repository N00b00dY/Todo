package main

import (
	"context"
	"distributor/dbs"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const gRpcPort = "5001"

type DBServer struct {
	dbs.UnimplementedDbServiceServer
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()
	dbs.RegisterDbServiceServer(server, &DBServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
