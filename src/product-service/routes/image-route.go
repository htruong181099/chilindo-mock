package routes

import (
	"chilindo/src/product-service/controllers"
	"github.com/gin-gonic/gin"
)

type IImageRoute interface {
	SetRoute()
}

type ImageRoute struct {
	ImageController controllers.IImageController
	Router          *gin.Engine
}

func (i ImageRoute) SetRoute() {
	api := i.Router.Group("/api/images")
	{
		api.POST("/", i.ImageController.CreateImage)
	}
}

func NewImageRoute(imageController controllers.IImageController, router *gin.Engine) *ImageRoute {
	return &ImageRoute{ImageController: imageController, Router: router}
}
