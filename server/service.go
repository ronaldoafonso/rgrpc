package main

import (
	"context"
	"errors"
	pb "github.com/ronaldoafonso/productinfo/productpb"
)

type ProductInfoServer struct {
	products map[string]*pb.Product
}

func (s *ProductInfoServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	in.Id = "id1"
	s.products[in.Id] = in
	return &pb.ProductID{
		Value: in.Id,
	}, nil
}

func (s *ProductInfoServer) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	if product, ok := s.products[in.Value]; ok {
		return product, nil
	}
	return nil, errors.New("Product does not exist.")
}

func NewProductInfoServer() *ProductInfoServer {
	return &ProductInfoServer{
		products: make(map[string]*pb.Product),
	}
}
