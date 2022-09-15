package main

import (
	"log"
	"os"
	"strconv"

	"test-client/internal/api"
	"test-client/internal/api/xgrpc"
	"test-client/internal/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGRPCClient() xgrpc.TestServiceClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("GRPC\tdid not connect: %v\n", err)
	}
	return xgrpc.NewTestServiceClient(conn)
}

func RunHTTPServer(errc chan<- error, port int, GRPC xgrpc.TestServiceClient) {
	log.Println("HTTP\tserver started")
	defer log.Println("HTTP\tserver finished")

	server, err := api.NewServer(app.App{GRPC: GRPC}, api.Config{
		Host: "localhost",
		Port: port,
	})
	if err != nil {
		errc <- err
		return
	}
	errc <- server.Serve()
}

func main() {
	log.Println("Run program")

	port := 8080
	var err error
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println(err)
			port = 8080
		}
	}

	GRPC := CreateGRPCClient()

	errc := make(chan error)
	go RunHTTPServer(errc, port, GRPC)
	err = <-errc
	if err != nil {
		log.Println(err)
	}
}
