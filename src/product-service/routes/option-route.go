package routes

import (
	"chilindo/src/product-service/controllers"
	"github.com/gin-gonic/gin"
)

type IOptionRoute interface {
	SetRouter()
}

type OptionRoute struct {
	OptionController controllers.IOptionController
	Router           *gin.Engine
}

func (o OptionRoute) SetRouter() {
	api := o.Router.Group("api/options")
	{
		api.POST("/", o.OptionController.CreateOption)
	}

}

func NewOptionRoute(optionController controllers.IOptionController, router *gin.Engine) *OptionRoute {
	return &OptionRoute{OptionController: optionController, Router: router}
}
