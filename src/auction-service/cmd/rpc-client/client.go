package rpc_client

import (
	"chilindo/pkg/pb/product"
	"chilindo/pkg/ssl"
	"google.golang.org/grpc"
	"log"
)

const (
	productClientPort = "localhost:50052"
)

type IRPCClient interface {
	SetUpProductClient() product.ProductServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpProductClient() product.ProductServiceClient {
	var opts []grpc.DialOption
	creds, tlsErr := ssl.LoadTLSCredentials()

	if tlsErr != nil {
		log.Fatalf("Failed to load credentials: %v", tlsErr)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(productClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", productClientPort)
	productClient := product.NewProductServiceClient(conn)
	return productClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
