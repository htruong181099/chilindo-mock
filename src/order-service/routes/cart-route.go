package routes

import (
	"chilindo/src/order-service/controllers"
	"github.com/gin-gonic/gin"
)

type ICartRoute interface {
	SetRouter()
}

type CartRoute struct {
	CartController controllers.ICartController
	router         *gin.Engine
}

func NewCartRoute(cartController controllers.ICartController, router *gin.Engine) *CartRoute {
	return &CartRoute{CartController: cartController, router: router}
}

func (c CartRoute) SetRouter() {
	api := c.router.Group("/api/cart")
	{
		api.GET("/")
		api.POST("/")
	}
}
