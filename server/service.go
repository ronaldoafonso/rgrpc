package main

import (
	"context"
	"errors"
	pb "github.com/ronaldoafonso/productinfo/productpb"
	"io"
	"log"
	"strconv"
)

var (
	ids      int
	products map[string]*pb.Product
)

type ProductInfoServer struct {
	products map[string]*pb.Product
}

func (s *ProductInfoServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	s.products = NewProductInfoServer()
	ids++
	in.Id = "id" + strconv.Itoa(ids)
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

func (s *ProductInfoServer) AddAllProducts(stream pb.ProductInfo_AddAllProductsServer) error {
	productIDs := []*pb.ProductID{}
	for {
		if product, err := stream.Recv(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		} else {
			ids++
			product.Id = "id" + strconv.Itoa(ids)
			s.products[product.Id] = product
			productIDs = append(productIDs, &pb.ProductID{Value: product.Id})
		}
	}
	return stream.SendAndClose(&pb.ProductIDs{ProductIDs: productIDs})
}

func (s *ProductInfoServer) GetAllProducts(empty *pb.Empty, stream pb.ProductInfo_GetAllProductsServer) error {
	for key, value := range s.products {
		if err := stream.Send(value); err != nil {
			log.Printf("Error getting product: %v (%v).", key, err)
			return err
		}
	}
	return nil
}

func NewProductInfoServer() map[string]*pb.Product {
	if products == nil {
		products = make(map[string]*pb.Product)
	}
	return products
}
