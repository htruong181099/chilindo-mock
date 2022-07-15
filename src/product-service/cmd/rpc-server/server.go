package rpc_server

import (
	"chilindo/pkg/pb/product"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	addr     = ":50052"
	certFile = "pkg/ssl/server.crt"
	keyFile  = "pkg/ssl/server.pem"
)

type ProductServer struct {
	product.ProductServiceServer
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

	product.RegisterProductServiceServer(s, &ProductServer{})

	log.Printf("listening on %s\n", addr)
	return s.Serve(lis)
}

func (p *ProductServer) GetProduct(ctx context.Context, in *product.GetProductRequest) (*product.GetProductResponse, error) {
	//log.Printf("Login request: %v\n", in)
	//
	//res, err := a.AuthService.CheckIsAdmin(in)
	//
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	//}
	//
	//if res == nil {
	//	return nil, status.Errorf(codes.NotFound, "User not found")
	//}
	//
	//return res, nil
	return nil, nil
}
