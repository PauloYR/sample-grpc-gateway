package main

import (
	"context"
	"gatewaysample/generated/hello"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	name := request.Name
	response := &hello.HelloResponse{
		Message: "Hello " + name,
	}
	return response, nil
}

func main() {

	network := "0.0.0.0"
	port := ":5000"
	portgRPC := ":5001"
	address := network + port

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	log.Printf("Serving gRPC on connection, listen in %s", port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = hello.RegisterGreeterHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    portgRPC,
		Handler: mux,
	}

	log.Printf("Serving gRPC-Gateway on connection, listen in %s", portgRPC)
	log.Fatalln(gwServer.ListenAndServe())
}
