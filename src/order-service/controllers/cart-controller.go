package controllers

import "chilindo/src/order-service/services"

type ICartController interface {
}

type CartController struct {
	CartService services.ICartService
}

func NewCartController(cartService services.ICartService) *CartController {
	return &CartController{CartService: cartService}
}
