package main

import (
	rpcClient "chilindo/src/product-service/cmd/rpc-client"
	"chilindo/src/product-service/cmd/rpc-server"
	"chilindo/src/product-service/controllers"
	controllers2 "chilindo/src/product-service/controllers/admin-rpc"
	"chilindo/src/product-service/database"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/routes"
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	ginPort              = ":3030"
	grpcServerPort       = "localhost:50052"
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

	//setup client
	grpcClient := rpcClient.NewRPCClient()
	adminClient := grpcClient.SetUpAdminClient()

	r := router()
	productRepo := repository.NewProductRepository(database.Instance)
	productScv := services.NewProductService(productRepo)
	productCtr := controllers.NewProductController(productScv)
	adminSrvCtr := controllers2.NewAdminServiceController(adminClient)
	productRoute := routes.NewProductRoute(productCtr, r, adminSrvCtr)
	productRoute.SetRouter()

	//Serve Gin Server
	go func() {
		if err := r.Run(ginPort); err != nil {
			log.Println("Open port is fail")
			return
		}
	}()

	//Serve Grpc Server
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = rpc_server.RunGRPCServer(true, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
