package routes

import (
	"chilindo/src/auction-service/controllers"
	controllers2 "chilindo/src/auction-service/controllers/user-rpc"
	"github.com/gin-gonic/gin"
)

type IAuctionRoute interface {
	SetRouter()
}

type AuctionRoute struct {
	AuctionController controllers.IAuctionController
	Router            *gin.Engine
	UserAuthSrvCtr    controllers2.IUserAuthServiceController
}

func NewAuctionRoute(auctionController controllers.IAuctionController, router *gin.Engine, userAuthSrvCtr controllers2.IUserAuthServiceController) *AuctionRoute {
	return &AuctionRoute{AuctionController: auctionController, Router: router, UserAuthSrvCtr: userAuthSrvCtr}
}

func (a AuctionRoute) SetRouter() {
	api := a.Router.Group("/api/auction")
	{
		api.GET("/", a.AuctionController.GetAuctions)
		api.GET("/:auctionId", a.AuctionController.GetAuctionById)
		api.POST("/", a.UserAuthSrvCtr.CheckIsAdmin(), a.AuctionController.CreateAuction)
	}
}
