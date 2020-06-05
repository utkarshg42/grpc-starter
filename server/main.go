//server for hosting a grpc service
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/utkarshg42/grpc-starter/handler"

	"google.golang.org/grpc"
)

const ip = ""
const port = "8088"

var (
	grpcServerAddr string
	restServerAddr string
)

func init() {
	fmt.Println("initializing server")
	grpcServerAddr = ":8088"
	restServerAddr = ":8087"
}

func grpcServer() error {
	log.Println("initializing grpc server with address ", grpcServerAddr)
	lis, err := net.Listen("tcp", grpcServerAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s", grpcServerAddr)
	}
	s := handler.Server{}

	grpcServer := grpc.NewServer()

	handler.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to Serve on %s", grpcServerAddr)
	}

	return nil
}

func restServer() error {
	log.Println("initializing rest server with address ", restServerAddr)
	ct := context.Background()
	ct, cancel := context.WithCancel(ct)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := handler.RegisterPingHandlerFromEndpoint(ct, mux, grpcServerAddr, opts)

	if err != nil {
		log.Fatalf("could not register rest service %s", err)
	}

	http.ListenAndServe(restServerAddr, mux)

	return nil
}

func main() {

	go grpcServer()

	go restServer()

	select {}

}
