package main

import (
	"context"
	pb "github.com/ronaldoafonso/productinfo/productpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"reflect"
	"testing"
)

func init() {
	initGRPCServer()
}

func TestAddAndGetProduct(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v.", err)
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	tests := []struct {
		product *pb.Product
		id      *pb.ProductID
	}{
		{
			product: &pb.Product{
				Name:        "prod1",
				Description: "Product 1",
				Price:       1.00,
			},
			id: &pb.ProductID{
				Value: "id1",
			},
		},
		{
			product: &pb.Product{
				Name:        "prod2",
				Description: "Product 2",
				Price:       2.00,
			},
			id: &pb.ProductID{
				Value: "id2",
			},
		},
		{
			product: &pb.Product{
				Name:        "prod3",
				Description: "Product 3",
				Price:       3.00,
			},
			id: &pb.ProductID{
				Value: "id3",
			},
		},
	}

	for _, test := range tests {
		got, err := client.AddProduct(context.Background(), test.product)
		if err != nil {
			t.Errorf("AddProduct error: %v.", err)
		}
		if reflect.DeepEqual(got, test.id) {
			t.Errorf("AddProduct error. Want: %v, got: %v.", test.id, got)
		}
	}

	for _, test := range tests {
		got, err := client.GetProduct(context.Background(), test.id)
		if err != nil {
			t.Errorf("GetProduct error: %v.", err)
		}
		if reflect.DeepEqual(got, test.product) {
			t.Errorf("GetProduct error. Want: %v, got: %v.", test.product, got)
		}
	}
}

func TestAddAllAndGetAllProducts(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v.", err)
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	products := []*pb.Product{
		&pb.Product{
			Name:        "prod1",
			Description: "Product 1",
			Price:       1.00,
		},
		&pb.Product{
			Name:        "prod2",
			Description: "Product 2",
			Price:       2.00,
		},
		&pb.Product{
			Name:        "prod3",
			Description: "Product 3",
			Price:       3.00,
		},
	}

	if stream, err := client.AddAllProducts(context.Background()); err != nil {
		t.Errorf("AddAllProducts error: %v.", err)
	} else {
		for _, product := range products {
			if err := stream.Send(product); err != nil {
				t.Errorf("AddAllProducts send error: %v.", err)
			}
		}
	}

	ids := []*pb.ProductID{
		&pb.ProductID{
			Value: "id1",
		},
		&pb.ProductID{
			Value: "id2",
		},
		&pb.ProductID{
			Value: "id3",
		},
	}

	stream, err := client.GetAllProducts(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("GetAllProducts error: %v.", err)
	}
	for _, id := range ids {
		if productID, err := stream.Recv(); err != nil {
			t.Errorf("GetAllProducts error: %v.", err)
		} else {
			if reflect.DeepEqual(id, productID) {
				t.Errorf("GetAllProducts error: Want: %v, got: %v.",
					id, productID)
			}
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
