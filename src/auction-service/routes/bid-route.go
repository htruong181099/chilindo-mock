package routes

import (
	"chilindo/src/auction-service/controllers"
	controllers2 "chilindo/src/auction-service/controllers/user-rpc"
	"github.com/gin-gonic/gin"
)

type IBidRoute interface {
	SetRouter()
}

type BidRoute struct {
	BidController  controllers.IBidController
	Router         *gin.Engine
	UserAuthSrvCtr controllers2.IUserAuthServiceController
}

func NewBidRoute(bidController controllers.IBidController, router *gin.Engine, userAuthSrvCtr controllers2.IUserAuthServiceController) *BidRoute {
	return &BidRoute{
		BidController:  bidController,
		Router:         router,
		UserAuthSrvCtr: userAuthSrvCtr,
	}
}

func (b BidRoute) SetRouter() {
	auctionBidAPI := b.Router.Group("/api/auction")
	{
		auctionBidAPI.GET("/:auctionId/bid", b.BidController.GetBidsOfAuction)
		auctionBidAPI.POST("/:auctionId/bid", b.UserAuthSrvCtr.CheckIsAuth(), b.BidController.CreateBid)
	}

	bidAPI := b.Router.Group("/api/bid")
	{
		bidAPI.GET("/:bidId", b.BidController.GetBidById)
	}
}
