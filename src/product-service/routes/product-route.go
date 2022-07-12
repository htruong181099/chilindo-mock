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
		api.GET("/:productId", p.ProductController.GetProductById)
		api.POST("/:productId", p.ProductController.UpdateProduct)
		api.DELETE("/:productId", p.ProductController.DeleteProduct)
		api.GET("/:productId/options", p.ProductController.GetOptions)
		api.POST("/:productId/options", p.ProductController.CreateOption)

	}

	optionAPI := p.Router.Group("/api/options")
	{
		optionAPI.GET("/:optionId", p.ProductController.GetOptionById)
		optionAPI.PATCH("/:optionId", p.ProductController.UpdateOption)
	}
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router}
}
