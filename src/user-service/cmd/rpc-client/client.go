package rpc_client

import (
	"chilindo/pkg/pb/auction"
	"chilindo/pkg/ssl"
	"google.golang.org/grpc"
	"log"
)

const (
	auctionClientPort = "localhost:50054"
)

type IRPCClient interface {
	SetUpClientAuction() auction.AuctionServiceServer
}
type RPCClient struct {
}

func (R RPCClient) SetUpClientAuction() auction.AuctionServiceClient {
	var opts []grpc.DialOption
	creds, tlsErr := ssl.LoadTLSCredentials()
	if tlsErr != nil {
		log.Fatalf("Failed to load credentials: %v", tlsErr)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(auctionClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", auctionClientPort)
	auctionClient := auction.NewAuctionServiceClient(conn)
	return auctionClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
