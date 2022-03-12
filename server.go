package main

import (
	"fmt"
	"log"
	"net"

	"gRPC/chat"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Server Start")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}