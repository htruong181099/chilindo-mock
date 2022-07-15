package routes

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/product-service/cmd/client"
	"chilindo/src/product-service/controllers"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	SetRouter()
}

type ProductRoute struct {
	ProductController controllers.IProductController
	Router            *gin.Engine
	AdminClient       admin.AdminServiceClient
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine, adminClient admin.AdminServiceClient) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router, AdminClient: adminClient}
}

func (p ProductRoute) SetRouter() {
	api := p.Router.Group("/api/products")
	{
		api.POST("/", client.CheckIsAdmin(p.AdminClient), p.ProductController.CreateProduct)
		api.GET("/", p.ProductController.GetProducts)
		api.GET("/:productId", p.ProductController.GetProductById)
		api.PUT("/:productId", client.CheckIsAdmin(p.AdminClient), p.ProductController.UpdateProduct)
		api.DELETE("/:productId", client.CheckIsAdmin(p.AdminClient), p.ProductController.DeleteProduct)
		api.GET("/:productId/options", p.ProductController.GetOptions)
		api.POST("/:productId/options", client.CheckIsAdmin(p.AdminClient), p.ProductController.CreateOption)

	}

	optionAPI := p.Router.Group("/api/options")
	{
		optionAPI.GET("/:optionId", p.ProductController.GetOptionById)
		optionAPI.PATCH("/:optionId", p.ProductController.UpdateOption)
		optionAPI.DELETE("/:optionId", p.ProductController.DeleteOption)
	}
}

//func NewProductRoute(productController controllers.IProductController, router *gin.Engine) *ProductRoute {
//	return &ProductRoute{ProductController: productController, Router: router}
//}
