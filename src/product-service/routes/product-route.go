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
	api := p.Router.Group("/admin")
	{
		api.POST("/product", p.ProductController.CreateProduct)
	}
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router}
}
