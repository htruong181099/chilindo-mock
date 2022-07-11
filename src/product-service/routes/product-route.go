package routes

import (
	"chilindo/src/product-service/controllers"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	SetRouter()
}

type ProductRoute struct {
	ProductController controllers.IProductController
	Router            *gin.Engine
}

func (p ProductRoute) SetRouter() {
	api := p.Router.Group("/api/products")
	{
		api.POST("/", p.ProductController.CreateProduct)
		api.GET("/", p.ProductController.GetProducts)
		api.GET("/:id", p.ProductController.GetProductById)
		api.POST("/:id", p.ProductController.UpdateProduct)
		api.DELETE("/:id", p.ProductController.DeleteProduct)
	}
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router}
}
