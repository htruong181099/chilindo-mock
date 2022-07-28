package rpc_client

import (
	"chilindo/pkg/pb/admin"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	adminClientPort = ":50051"
)

type IRPCClient interface {
	SetUpAdminClient() admin.AdminServiceClient
}

type RPCClient struct{}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}

func (R RPCClient) SetUpAdminClient() admin.AdminServiceClient {
	//var opts []grpc.DialOption
	//creds, err := ssl.LoadTLSCredentials()
	//if err != nil {
	//	log.Fatalf("Failed to load credentials: %v", err)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))

	addr := os.Getenv("USER_SRV_HOST")

	conn, dialErr := grpc.Dial(addr, grpc.WithInsecure())
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", dialErr)
	}

	adminClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to AdminService on port ", adminClientPort)
	return adminClient
}
