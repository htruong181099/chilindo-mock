package routes

import (
	"chilindo/src/product-service/controllers"
	controllers2 "chilindo/src/product-service/controllers/admin-rpc"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	SetRouter()
}

type ProductRoute struct {
	ProductController  controllers.IProductController
	Router             *gin.Engine
	AdminSrvController controllers2.IAdminServiceController
}

func NewProductRoute(productController controllers.IProductController, router *gin.Engine, adminSrvController controllers2.IAdminServiceController) *ProductRoute {
	return &ProductRoute{
		ProductController:  productController,
		Router:             router,
		AdminSrvController: adminSrvController,
	}
}

func (p ProductRoute) SetRouter() {
	api := p.Router.Group("/api/products")
	{
		api.POST("/", p.AdminSrvController.CheckIsAdmin(), p.ProductController.CreateProduct)
		api.GET("/", p.ProductController.GetProducts)
		api.GET("/:productId", p.ProductController.GetProductById)
		api.PUT("/:productId", p.AdminSrvController.CheckIsAdmin(), p.ProductController.UpdateProduct)
		api.DELETE("/:productId", p.AdminSrvController.CheckIsAdmin(), p.ProductController.DeleteProduct)
		api.GET("/:productId/options", p.ProductController.GetOptions)
		api.POST("/:productId/options", p.AdminSrvController.CheckIsAdmin(), p.ProductController.CreateOption)

	}

	optionAPI := p.Router.Group("/api/options")
	{
		optionAPI.GET("/:optionId", p.ProductController.GetOptionById)
		optionAPI.PATCH("/:optionId", p.AdminSrvController.CheckIsAdmin(), p.ProductController.UpdateOption)
		optionAPI.DELETE("/:optionId", p.AdminSrvController.CheckIsAdmin(), p.ProductController.DeleteOption)
	}
}
