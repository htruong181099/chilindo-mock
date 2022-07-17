package rpc_server

import (
	"chilindo/pkg/pb/auction"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/services"
	"chilindo/src/user-service/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	grpcServerPort = ":50053"
	certFile       = "pkg/ssl/server.crt"
	keyFile        = "pkg/ssl/server.pem"
)

type auctionServer struct {
	auction.AuctionServiceServer
	AuctionService services.IAuctionService
}

func RunGRPCServerAuction(enabledTLS bool, lis net.Listener) error {
	var opts []grpc.ServerOption
	if enabledTLS {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Println("RunGRPCServerAuction: Error credentials NewServer", err)
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)

	repoAuction := repository.NewAuctionRepository(database.Instance)
	srvAuction := services.NewAuctionService(repoAuction)

	auction.RegisterAuctionServiceServer(s, &auctionServer{
		AuctionService: srvAuction,
	})
	log.Printf("Server AuctionService server on %s\n", grpcServerPort)
	return s.Serve(lis)
}
