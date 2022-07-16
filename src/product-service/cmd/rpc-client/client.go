package rpc_client

import (
	"chilindo/pkg/pb/admin"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	grpcClientPort = "localhost:50051"
	certFile       = "pkg/ssl/ca.crt"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, err
	}
	return creds, nil
}
func SetupAdminClient() admin.AdminServiceClient {
	var opts []grpc.DialOption
	creds, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, dialErr := grpc.Dial(grpcClientPort, opts...)
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	adminClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to AdminService on port ", grpcClientPort)
	return adminClient
}
