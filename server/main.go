package main

import (
	"fmt"
	"context"
	"time"
	"log"
	"net"
	
	demo_grpc "grpc_server/proto"
	"google.golang.org/grpc"  // go grpc package. `go get -u google.golang.org/grpc`
)

// The server strcut/object must implement the interface defined in the proto file
type Server struct {
	// placeholder for all RPC methods define in .proto file. Safeguard if we miss any methods, but still compile
	// The method for the following strcut will also be the method in the outer struct
	// golang strcut embedding. Anonmous embedded Fields
	// The method will be "override"
	demo_grpc.UnimplementedDemoServiceServer  
}

/*
method for the strcut Server
SayHello is the RPC name
It takes a context and a pointer to a HelloRequest object, and outputs a pointer to a HelloResponse object and an error
Rest is implementation
*/

func (s *Server) SayHello(ctx context.Context, in *demo_grpc.HelloRequest) (*demo_grpc.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &demo_grpc.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) Adder(ctx context.Context, in *demo_grpc.AdderRequest) (*demo_grpc.AdderResponse, error) {
	log.Printf("Received: %v, %v", in.GetA(), in.GetB())
	time.Sleep(2 * time.Second)

	return &demo_grpc.AdderResponse{Result: in.GetA() + in.GetB()}, nil
}

func main() {
	fmt.Println("Running server/main.go")

	port := 10000
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))  // socket listen
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()
	// The 2nd argument is the server object that implements the interface defined in the proto file
	demo_grpc.RegisterDemoServiceServer(grpcServer, &Server{})  // create a new instant of the server struct and register it to the grpc server

	fmt.Printf("Start listening on port %v\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	   }
}
