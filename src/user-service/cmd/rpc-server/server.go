package rpc_server

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/user-service/database"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	addr     = ":50051"
	certFile = "src/product-service/cmd/ssl/server.crt"
	keyFile  = "src/product-service/cmd/ssl/server.pem"
)

type AdminServer struct {
	admin.AdminServiceServer
	AuthService services.IAuthService
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

	userRepo := repository.NewUserRepository(database.Instance)
	AuthService := services.NewAuthService(userRepo)

	admin.RegisterAdminServiceServer(s, &AdminServer{
		AuthService: AuthService,
	})

	log.Printf("listening on %s\n", addr)
	return s.Serve(lis)
}

func (a *AdminServer) CheckIsAdmin(ctx context.Context, in *admin.CheckIsAdminRequest) (*admin.CheckIsAdminResponse, error) {
	log.Printf("Login request: %v\n", in)

	res, err := a.AuthService.CheckIsAdmin(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	if res == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return res, nil
}
