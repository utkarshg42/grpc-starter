package main

import (
	"context"
	"fmt"
	"log"

	"github.com/utkarshg42/grpc-starter/handler"
	"google.golang.org/grpc"
)

var serverAddr string

func init() {
	fmt.Println("initializing client")
	serverAddr = ":8088"
}

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error in connecting to the server at %s", serverAddr)
	}
	defer conn.Close()

	cl := handler.NewPingClient(conn)

	response, err := cl.SayHello(context.Background(), &handler.PingMessage{Greeting: "Hi server"})

	if err != nil {
		log.Fatalf("error in greeting to the server %s", err)
	}
	log.Printf("Response received %s from server", response)
}
