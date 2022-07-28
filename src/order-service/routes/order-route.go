package routes

import (
	"chilindo/src/order-service/controllers"
	"github.com/gin-gonic/gin"
)

type IOrderRoute interface {
	SetRouter()
}

type OrderRoute struct {
	OrderController controllers.IOrderController
	router          *gin.Engine
}

func NewOrderRoute(orderController controllers.IOrderController, router *gin.Engine) *OrderRoute {
	return &OrderRoute{OrderController: orderController, router: router}
}

func (o OrderRoute) SetRouter() {
	api := o.router.Group("/api/order")
	{
		api.GET("/")
		api.POST("/")
	}
}
