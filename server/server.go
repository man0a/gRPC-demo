package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gprcdemo.com/todo/todo"
)

const (
	port = ":14586"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}
	grpcServer := grpc.NewServer()
	log.Printf("Registering Server")
	// attach the Ping service to the server
	todo.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
