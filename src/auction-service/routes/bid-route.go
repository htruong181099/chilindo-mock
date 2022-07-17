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
	api := b.Router.Group("/api/bid")
	{
		api.GET("/:bidId", b.BidController.GetBidById)
		api.GET("/:auctionId", b.BidController.GetBidsOfAuction)
		api.POST("/:auctionId", b.UserAuthSrvCtr.CheckIsAuth(), b.BidController.CreateBid)
	}
}
