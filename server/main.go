package main

import (
	pb "github.com/ronaldoafonso/productinfo/productpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v.", err)
	}

	server := grpc.NewServer()
	pb.RegisterProductInfoServer(server, NewProductInfoServer())

	log.Println("Starting Product Info Server.")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v.", err)
	}
}
