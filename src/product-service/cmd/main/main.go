package main

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/product-service/cmd/rpc-server"
	"chilindo/src/product-service/controllers"
	controllers2 "chilindo/src/product-service/controllers/admin-rpc"
	"chilindo/src/product-service/database"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/routes"
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	ginPort              = ":3030"
	grpcClientPort       = "localhost:50051"
	grpcServerPort       = "localhost:50052"
	certFile             = "pkg/ssl/ca.crt"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, err
	}
	return creds, nil
}

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
	var opts []grpc.DialOption
	creds, err := loadTLSCredentials()

	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	//grpc CLient
	conn, err := grpc.Dial(grpcClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	adminClient := admin.NewAdminServiceClient(conn)

	r := router()
	//DI Product
	productRepo := repository.NewProductRepository(database.Instance)
	productScv := services.NewProductService(productRepo)
	productCtr := controllers.NewProductController(productScv)
	adminSrvCtr := controllers2.NewAdminServiceController()
	productRoute := routes.NewProductRoute(productCtr, r, adminSrvCtr, adminClient)
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
	log.Println("gRPC server admin is running")

}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
