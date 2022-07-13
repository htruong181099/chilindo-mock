package main

import (
	"chilindo/src/auction-service/controllers"
	"chilindo/src/auction-service/database"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/routes"
	"chilindo/src/auction-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
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

	auctionRepository := repository.NewAuctionRepository(database.Instance)
	auctionService := services.NewAuctionService(auctionRepository)
	auctionController := controllers.NewAuctionController(auctionService)
	auctionRouter := routes.NewAuctionRoute(auctionController, r)
	auctionRouter.SetRouter()

	if err := r.Run(":3030"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
