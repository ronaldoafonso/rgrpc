package main

import (
	"context"
	pb "github.com/ronaldoafonso/productinfo/productpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

func TestNewProductInfoServer(t *testing.T) {
	s1 := NewProductInfoServer()
	if s1 == nil {
		t.Error("NewProductInforServer returning nil and not a singleton.")
	}
}

func TestAddProduct(t *testing.T) {
	initGRPCServer()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v.", err)
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	tests := []struct {
		name        string
		description string
		price       float32
		want        *pb.ProductID
	}{
		{"prod1", "Product 1", 1.00, &pb.ProductID{Value: "id1"}},
		{"prod2", "Product 2", 2.00, &pb.ProductID{Value: "id2"}},
		{"prod3", "Product 3", 3.00, &pb.ProductID{Value: "id3"}},
	}

	for _, test := range tests {
		product := &pb.Product{
			Name:        test.name,
			Description: test.description,
			Price:       test.price,
		}
		got, err := client.AddProduct(context.Background(), product)
		if err != nil {
			t.Errorf("AddProduct error: %v.", err)
		}
		if got.Value != test.want.Value {
			t.Errorf("AddProduct error. Want: %v, got: %v.",
				test.want.Value, got.Value)
		}
	}
}

func initGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v.", err)
	}
	server := grpc.NewServer()
	pb.RegisterProductInfoServer(server, &ProductInfoServer{})
	reflection.Register(server)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v.", err)
		}
	}()
}
