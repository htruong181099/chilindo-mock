package routes

import (
	"chilindo/src/auction-service/controllers"
	"github.com/gin-gonic/gin"
)

type IBidRoute interface {
	SetRouter()
}

type BidRoute struct {
	BidController controllers.IBidController
	Router        *gin.Engine
}

func (b BidRoute) SetRouter() {
	api := b.Router.Group("/api/bid")
	{
		api.GET("/")
		api.POST("/")
	}
}
