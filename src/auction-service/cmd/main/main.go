package main

import (
	rpcClient "chilindo/src/auction-service/cmd/rpc-client"
	rpc_server "chilindo/src/auction-service/cmd/rpc-server"
	"chilindo/src/auction-service/controllers"
	controllers2 "chilindo/src/auction-service/controllers/user-rpc"
	"chilindo/src/auction-service/database"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/routes"
	"chilindo/src/auction-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	ginPort              = ":3031"
	grpcServerPort       = ":50053"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	connectString := "root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true"
	if envErr == nil {
		connectString = os.Getenv(DB_CONNECTION_STRING)
	}
	database.Connect(connectString)
	database.Migrate()

	//Create new gRPC Client

	grpcClient := rpcClient.NewRPCClient()
	productClient := grpcClient.SetUpProductClient()

	r := router()
	auctionRepository := repository.NewAuctionRepository(database.Instance)
	auctionService := services.NewAuctionService(auctionRepository)
	auctionController := controllers.NewAuctionController(auctionService, productClient)
	auctionRouter := routes.NewAuctionRoute(auctionController, r)
	auctionRouter.SetRouter()

	//DI Bid
	authClient := grpcClient.SetUpAuthClient()
	bidRepo := repository.NewBidRepository(database.Instance)
	bidSrv := services.NewBidService(bidRepo)
	bidCtr := controllers.NewBidController(bidSrv)
	authCtr := controllers2.NewUserAuthServiceController(authClient)
	bidRoute := routes.NewBidRoute(bidCtr, r, authCtr)
	bidRoute.SetRouter()
	//Server Gin run
	go func() {
		if err := r.Run(ginPort); err != nil {
			log.Println("Open port is fail")
			return
		}
	}()
	//Serve Grpc run
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err = rpc_server.RunGRPCServerAuction(true, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
