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

//git checkout -b features/product
//git pull origin features/product ->
//git add ./src/product-service
//git commit -m ""
//git push origin features/product

func main() {

	database.Connect("root:@Duy123456789@tcp(localhost:3306)/chilindo?parseTime=true")
	database.Migrate()

	r := router()
	//DI Product
	productRepo := repository.NewProductRepository(database.Instance)
	productScv := services.NewProductService(productRepo)
	productCtr := controllers.NewProductController(productScv)
	productRoute := routes.NewProductRoute(productCtr, r)
	productRoute.SetRouter()

	if err := r.Run(":3030"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
