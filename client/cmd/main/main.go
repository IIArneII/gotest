package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"test-client/internal/api/xgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Printf("Start program\n")
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	client := xgrpc.NewTestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	msg, err := client.SendMessage(ctx, &xgrpc.Message{
		Msg: "Невероятно",
	})
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	} else {
		fmt.Printf("Received a response: %s\n", msg.Msg)
	}
}
