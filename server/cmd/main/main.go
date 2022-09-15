package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"test-server/internal/api"
	"test-server/internal/api/xgrpc"
	"test-server/internal/app"
)

func RunGRPCServer(errc chan<- error, app *app.App) {
	log.Println("GRPC\tserver started")
	defer log.Println("GRPC\tserver finished")

	server := grpc.NewServer()

	xgrpc.RegisterTestServiceServer(server, xgrpc.New(app))

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		errc <- err
		return
	}

	errc <- server.Serve(l)
}

func RunHTTPServer(errc chan<- error, app *app.App) {
	log.Println("HTTP\tserver started")
	defer log.Println("HTTP\tserver finished")

	server, err := api.NewServer(app, api.Config{
		Host: "localhost",
		Port: 8081,
	})
	if err != nil {
		errc <- err
		return
	}
	errc <- server.Serve()
}

func main() {
	log.Println("Run program")

	app := &app.App{
		Clients: make(map[string]app.Client),
	}

	errc := make(chan error)
	go RunGRPCServer(errc, app)
	go RunHTTPServer(errc, app)
	err := <-errc
	if err != nil {
		log.Println(err)
	}
}
