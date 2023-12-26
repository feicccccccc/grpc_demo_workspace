package main

import (
	"context"
	"fmt"
	"log"

	demo_grpc "grpc_server/proto" // calling from the server proto package

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Running client/main.go")
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := demo_grpc.NewDemoServiceClient(conn)

	// SayHello
	req := demo_grpc.HelloRequest{Name: "gRPC"}
	// req.ProtoMessage() // this is marker method for protobuf. Empty method, used to prevent interface from being implemented outside of the proto package

	res, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("SayHello Response: %v", res.GetMessage())

	// Adder
	req2 := demo_grpc.AdderRequest{A: 1, B: 2}
	res2, err := client.Adder(context.Background(), &req2)
	if err != nil {
		log.Fatalf("failed to call Adder: %v", err)
	}
	log.Printf("Adder Response: %v", res2.GetResult())
}
