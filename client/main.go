package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

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

	// StringToChar
	req3 := demo_grpc.HelloRequest{Name: "send"}

	stream, err := client.StringToChar(context.Background(), &req3)
	if err != nil {
		log.Fatalf("failed to call StringToChar: %v", err)
	}
	
	// forever loop
	for {
		msg, err := stream.Recv()

		// end of stream
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}

		log.Printf("StringToChar Response: %c", rune(msg.GetChar()))
	}

	// CharToString
	chars := []uint32{'g', 'R', 'P', 'C'}

	stream2, err := client.CharToString(context.Background())
	if err != nil {
		log.Fatalf("failed to call CharToString: %v", err)
	}

	for _, c := range chars {
		log.Printf("Sent: %c", rune(c))
		time.Sleep(500 * time.Millisecond)
		if err := stream2.Send(&demo_grpc.CharRequest{Char: c}); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
	}

	res3, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}
	log.Printf("CharToString Response: %v", res3)

	// AllCharUpper
	chars2 := []uint32{'l', 'o', 'w', 'e', 'r'}
	stream3, err := client.AllCharUpper(context.Background())
	if err != nil {
		log.Fatalf("failed to call AllCharUpper: %v", err)
	}

	// Force the goroutine to wait for the server to finish sending data
	var wg sync.WaitGroup

	// a new goroutine to recieved data from the server
	wg.Add(1) // add 1 to wg
	go func() {
		for {
			in, err := stream3.Recv()
			if err == io.EOF {
				wg.Done()  // signal wg that we are done
				return
			}
			if err != nil {
				log.Fatalf("failed to receive: %v", err)
			}
			log.Printf("AllCharUpper Response: %c", rune(in.GetChar()))
		}
	}()
	
	// send data to the server
	for _, c := range chars2 {
		log.Printf("Sent: %c", rune(c))
		time.Sleep(200 * time.Millisecond)
		if err := stream3.Send(&demo_grpc.CharRequest{Char: c}); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
	}
	stream3.CloseSend()  // trigger EOF
	wg.Wait()  // wait till goroutine is done
}
