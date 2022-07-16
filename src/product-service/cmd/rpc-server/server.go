package rpc_server

import (
	"chilindo/pkg/pb/product"
	"chilindo/src/product-service/database"
	"chilindo/src/product-service/dtos"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	grpcServerPort = ":50052"
	certFile       = "pkg/ssl/server.crt"
	keyFile        = "pkg/ssl/server.pem"
)

type ProductServer struct {
	product.ProductServiceServer
	productService services.IProductService
}

func RunGRPCServer(enabledTLS bool, lis net.Listener) error {
	var opts []grpc.ServerOption
	if enabledTLS {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	fmt.Println("Inc", database.Instance)
	repoProduct := repository.NewProductRepository(database.Instance)
	serviceProduct := services.NewProductService(repoProduct)

	product.RegisterProductServiceServer(s, &ProductServer{
		productService: serviceProduct,
	})

	log.Printf("Serve ProductService server on %s\n", grpcServerPort)
	return s.Serve(lis)
}

func (p *ProductServer) GetProduct(ctx context.Context, in *product.GetProductRequest) (*product.GetProductResponse, error) {
	log.Printf("Login request: %v\n", in)
	pid := in.GetProductId()
	if pid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "InvalidArgument productId= %v", pid)
	}
	var dto dtos.ProductDTO
	dto.ProductId = in.GetProductId()
	prod, err := p.productService.GetProductById(&dto)
	fmt.Println("Check prod", prod)
	log.Println("Check error", err)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}
	if prod == nil {
		return &product.GetProductResponse{
			IsFound: false,
		}, nil
	}
	response := &product.GetProductResponse{
		IsFound:     true,
		Id:          prod.Id,
		Name:        prod.Name,
		Price:       prod.Price,
		Description: prod.Description,
		Quantity:    int32(prod.Quantity),
	}
	return response, nil
}
