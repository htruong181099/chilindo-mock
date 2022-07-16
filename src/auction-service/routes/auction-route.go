package routes

import (
	"chilindo/src/auction-service/controllers"
	"github.com/gin-gonic/gin"
)

type IAuctionRoute interface {
	SetRouter()
}

type AuctionRoute struct {
	AuctionController controllers.IAuctionController
	Router            *gin.Engine
}

func NewAuctionRoute(auctionController controllers.IAuctionController, router *gin.Engine) *AuctionRoute {
	return &AuctionRoute{AuctionController: auctionController, Router: router}
}

func (a AuctionRoute) SetRouter() {
	api := a.Router.Group("/api/auction")
	{
		api.GET("/")
		api.POST("/", a.AuctionController.CreateAuction)
	}
}
