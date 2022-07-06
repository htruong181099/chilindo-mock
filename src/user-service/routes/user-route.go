package routes

import (
	"chilindo/controllers"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	SetRouter()
}
type UserRoute struct {
	UserController controllers.IUserController
	Router         *gin.Engine
}

func (u UserRoute) SetRouter() {
	api := u.Router.Group("/api/user")
	{
		api.GET("/address/:addressId", u.UserController.GetAddress)
	}

}

func NewUserRoute(userController controllers.IUserController, router *gin.Engine) *UserRoute {
	return &UserRoute{UserController: userController, Router: router}
}
