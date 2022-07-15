package main

import (
	rpc_server "chilindo/src/user-service/cmd/rpc-server"
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
	addr = ":50051"
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

	//DI Auth
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
		if err := r.Run(":3000"); err != nil {
			log.Println("Open port is fail")
			return
		}
		log.Println("Run port 3000")
	}()

	lis, err := net.Listen("tcp", addr)
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
