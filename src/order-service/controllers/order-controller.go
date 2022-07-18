package controllers

import "chilindo/src/order-service/services"

type IOrderController interface {
}

type OrderController struct {
	OrderService services.IOrderService
}

func NewOrderController(orderService services.IOrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}
