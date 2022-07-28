package controllers

import (
	"chilindo/src/order-service/dtos"
	"chilindo/src/order-service/models"
	"chilindo/src/order-service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ICartController interface {
	GetCart(ctx *gin.Context)
	GetListCarts(ctx *gin.Context)
}

type CartController struct {
	CartService services.ICartService
}

func (c CartController) GetCart(ctx *gin.Context) {
	var cartBody *models.Cart
	if err := ctx.ShouldBindJSON(&cartBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": " fail to get cart ",
		})
		log.Println("GetCart: Error ShouldBindJSON in pgk controller", err)
		ctx.Abort()
		return
	}
	var dto dtos.CartDTO
	dto.CartId = cartBody.BidId
	cart, err := c.CartService.GetCart(&dto)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": " Not found cart ",
		})
		log.Println("GetCart: Error to call CartService  in pgk controller", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

func (c CartController) GetListCarts(ctx *gin.Context) {
	var cartBody *models.Cart
	if err := ctx.ShouldBindJSON(&cartBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": " fail to get cart ",
		})
		log.Println("GetListCarts: Error ShouldBindJSON in pgk controller", err)
		ctx.Abort()
		return
	}
	var dto dtos.CartDTO
	dto.CartId = cartBody.UserId
	cart, err := c.CartService.GetCart(&dto)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": " Not found cart",
		})
		log.Println("GetListCarts: Error to call CartService  in pgk controller", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

func NewCartController(cartService services.ICartService) *CartController {
	return &CartController{CartService: cartService}
}
