package rpc_client

import (
	"chilindo/pkg/pb/admin"
	"chilindo/pkg/ssl"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	adminClientPort = "localhost:50051"
)

type IRPCClient interface {
	SetUpAdminClient() admin.AdminServiceClient
}

type RPCClient struct{}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}

func (R RPCClient) SetUpAdminClient() admin.AdminServiceClient {
	var opts []grpc.DialOption
	creds, err := ssl.LoadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, dialErr := grpc.Dial(adminClientPort, opts...)
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	adminClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to AdminService on port ", adminClientPort)
	return adminClient
}
