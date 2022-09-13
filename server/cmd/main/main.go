package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"gtest/server/internal/api/xgrpc"
)

func main() {
	fmt.Printf("Start program\n")
	s := grpc.NewServer()

	xgrpc.RegisterTestServiceServer(s, xgrpc.New())

	fmt.Printf("Create listener\n")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Run server\n")
	err = s.Serve(l)
	if err != nil {
		log.Fatalln(err)
	}
}
