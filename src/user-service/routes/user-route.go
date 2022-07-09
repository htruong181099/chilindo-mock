package routes

import (
	"chilindo/src/user-service/controllers"
	"chilindo/src/user-service/utils"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	SetRouter()
}
type UserRoute struct {
	UserController controllers.IUserController
	Router         *gin.Engine
	MW             *utils.SMiddleWare
}

func (u UserRoute) SetRouter() {

	api := u.Router.Group("/api/user").Use(u.MW.MiddleWare())
	{
		api.GET("/address", u.UserController.GetAddress)
		api.GET("/address/:addressId", u.UserController.GetAddressById)
		api.POST("/address/", u.UserController.CreateAddressByUserId)
		//api.DELETE("/address/", u.UserController.)
	}

}

func NewUserRoute(userController controllers.IUserController, router *gin.Engine) *UserRoute {
	return &UserRoute{UserController: userController, Router: router}
}
