package main

import (
	"chilindo/src/product-service/controllers"
	"chilindo/src/product-service/database"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/routes"
	"chilindo/src/product-service/services"
	"github.com/gin-gonic/gin"
	"log"
)

//const (
//	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
//)

func main() {
	//envErr := godotenv.Load(".env")
	//if envErr != nil {
	//	log.Fatalf("Error loading .env file")
	//}

	//connectString := "root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true"
	//if envErr == nil {
	//	connectString = os.Getenv(DB_CONNECTION_STRING)
	//}
	database.Connect("root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true")
	database.Migrate()

	r := router()
	//DI Product
	productRepo := repository.NewProductRepository(database.Instance)
	productScv := services.NewProductService(productRepo)
	productCtr := controllers.NewProductController(productScv)
	productRoute := routes.NewProductRoute(productCtr, r)
	productRoute.SetRouter()
	//
	////DI Image
	//imageRepo := repository.NewImageRepository(database.Instance)
	//imageSvc := services.NewImageService(imageRepo)
	//imageCtr := controllers.NewImageController(imageSvc)
	//imageRoutes := routes.NewImageRoute(imageCtr, r)
	//imageRoutes.SetRoute()

	//Di Option
	//optionRepo := repository.NewOptionRepository(database.Instance)
	//optionSvc := services.NewOptionService(optionRepo)
	//optionCtr := controllers.NewOptionController(optionSvc)
	//optionRoutes := routes.NewOptionRoute(optionCtr, r)
	//optionRoutes.SetRouter()

	if err := r.Run(":3030"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
