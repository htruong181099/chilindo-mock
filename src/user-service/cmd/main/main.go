package main

import (
	rpcServer "chilindo/src/user-service/cmd/rpc-server"
	"chilindo/src/user-service/controllers"
	"chilindo/src/user-service/database"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/routes"
	"chilindo/src/user-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
)

const (
	ginPort = ":3000"
	addr    = ":50051"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Println("Error loading .env file")
	}
	connectString := os.Getenv(DB_CONNECTION_STRING)

	if envErr == nil {
		connectString = os.Getenv(DB_CONNECTION_STRING)
	}

	database.Connect(connectString)
	database.Migrate()

	r := router()

	userRepo := repository.NewUserRepository(database.Instance)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)
	authRouter := routes.NewAuthRoute(authController, r)
	authRouter.SetRouter()

	addressRepo := repository.NewAddressRepository(database.Instance)
	userService := services.NewUserService(userRepo, addressRepo)
	userController := controllers.NewUserController(userService)
	userRouter := routes.NewUserRoute(userController, r)
	userRouter.SetRouter()

	go func() {
		if err := r.Run(ginPort); err != nil {
			log.Println("Open port is fail")
			return
		}
		log.Println("Run port 3000")
	}()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = rpcServer.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server admin is running")

}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
