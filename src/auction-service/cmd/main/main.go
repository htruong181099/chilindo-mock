package main

import (
	"chilindo/pkg/pb/product"
	"chilindo/src/auction-service/controllers"
	"chilindo/src/auction-service/database"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/routes"
	"chilindo/src/auction-service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	prodClientPort       = "localhost:50052"
	ginPort              = ":3031"
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

	conn := setUpProductClient()

	productClient := product.NewProductServiceClient(conn)
	fmt.Println(productClient)

	r := router()
	//DI Auth
	auctionRepository := repository.NewAuctionRepository(database.Instance)
	auctionService := services.NewAuctionService(auctionRepository)
	auctionController := controllers.NewAuctionController(auctionService)
	auctionRouter := routes.NewAuctionRoute(auctionController, r)
	auctionRouter.SetRouter()

	if err := r.Run(ginPort); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func setUpProductClient() *grpc.ClientConn {
	var opts []grpc.DialOption
	creds, err := loadTLSCredentials()

	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(prodClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	log.Println("Listening from port :", prodClientPort)
	return conn
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
