//server for hosting a grpc service
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/utkarshg42/grpc-starter/handler"

	"google.golang.org/grpc"
)

var (
	grpcServerAddr    string
	restServerAddr    string
	metricsServerAddr string
)

func init() {
	fmt.Println("initializing server")
	grpcServerAddr = ":8088"
	restServerAddr = ":8087"
	metricsServerAddr = ":8089"
}

func grpcServer() error {
	log.Println("initializing grpc server with address ", grpcServerAddr)

	defer func() {
		fmt.Println("returned from grpc server")
	}()

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
	defer func() {
		fmt.Println("returned from rest server")
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := handler.RegisterPingHandlerFromEndpoint(ct, mux, grpcServerAddr, opts)

	if err != nil {
		log.Fatalf("could not register rest service %s", err)
	}
	http.ListenAndServe(restServerAddr, mux)

	return nil
}

func metricsServer() error {
	log.Println("initializing metrics server with address ", metricsServerAddr)
	http.ListenAndServe(metricsServerAddr, promhttp.Handler())
	return nil
}

func main() {

	go grpcServer()

	go restServer()

	go metricsServer()

	select {}

}
