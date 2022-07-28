package rpc_client

import (
	"chilindo/pkg/pb/admin"
	"chilindo/pkg/pb/product"
	"chilindo/pkg/ssl"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	authClientPort    = "localhost:50051"
	productClientPort = "localhost:50052"
)

type IRPCClient interface {
	SetUpProductClient() product.ProductServiceClient
	SetUpAdminClient() admin.AdminServiceClient
}

type RPCClient struct{}

func (r *RPCClient) SetUpAdminClient() admin.AdminServiceClient {
	var opts []grpc.DialOption
	creds, err := ssl.LoadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, dialErr := grpc.Dial(authClientPort, opts...)
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", dialErr)
	}

	authClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to AdminService on port ", authClientPort)
	return authClient
}

func (r *RPCClient) SetUpProductClient() product.ProductServiceClient {
	var opts []grpc.DialOption
	creds, tlsErr := ssl.LoadTLSCredentials()

	if tlsErr != nil {
		log.Fatalf("Failed to load credentials: %v", tlsErr)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial(productClientPort, opts...)
	conn, err := grpc.Dial(productClientPort, grpc.WithInsecure())
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
