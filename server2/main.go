package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	port := 7901
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterShellServiceServer(grpcServer, &ShellService{})
	log.Printf("starting server. http port: %d ", port)
	grpcServer.Serve(lis)
}
