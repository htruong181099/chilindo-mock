package main

import (
	"chilindo/src/user-service/controllers"
	"chilindo/src/user-service/database"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/routes"
	"chilindo/src/user-service/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	database.Connect("root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true")
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

	if err := r.Run(":3000"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
