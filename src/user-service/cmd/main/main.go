package main

import (
	"chilindo/controllers"
	"chilindo/database"
	"chilindo/repository"
	"chilindo/routes"
	"chilindo/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Connect("root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true")
	database.Migrate()

	//DI Auth
	r := router()

	userRepo := repository.NewUserRepository(database.Instance)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	authRouter := routes.NewAuthRoute(userController, r)
	authRouter.SetRouter()

	if err := r.Run(":3000"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
