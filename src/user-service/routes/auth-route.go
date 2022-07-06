package routes

import (
	"chilindo/controllers"
	"github.com/gin-gonic/gin"
)

type IAuthRoute interface {
	SetRouter()
}
type AuthRoute struct {
	UserController controllers.IUserController
	Router         *gin.Engine
}

func (u AuthRoute) SetRouter() {
	api := u.Router.Group("/api/auth")
	{
		api.POST("/signup", u.UserController.SignUp)
		api.POST("/signin", u.UserController.SignIn)
	}

}

func NewAuthRoute(userController controllers.IUserController, router *gin.Engine) *AuthRoute {
	return &AuthRoute{UserController: userController, Router: router}
}
