package main

import (
	"chilindo/src/order-service/controllers"
	"chilindo/src/order-service/database"
	"chilindo/src/order-service/repository"
	"chilindo/src/order-service/routes"
	"chilindo/src/order-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
)

const (
	ginPort = ":3032"
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

	r := router()

	cartRepo := repository.NewCartRepository(database.Instance)
	cartService := services.NewCartService(cartRepo)
	cartController := controllers.NewCartController(cartService)
	cartRouter := routes.NewCartRoute(cartController, r)
	cartRouter.SetRouter()

	orderRepo := repository.NewOrderRepository(database.Instance)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)
	orderRouter := routes.NewOrderRoute(orderController, r)
	orderRouter.SetRouter()

	if err := r.Run(ginPort); err != nil {
		log.Fatalln("Open port is fail")
		return
	}
	log.Printf("Gin server serve on port %v /n", ginPort)

}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
