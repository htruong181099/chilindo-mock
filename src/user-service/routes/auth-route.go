package routes

import (
	"chilindo/controllers"
	"github.com/gin-gonic/gin"
)

type IAuthRoute interface {
	SetRouter()
}
type AuthRoute struct {
	AuthController controllers.IAuthController
	Router         *gin.Engine
}

func (u AuthRoute) SetRouter() {
	api := u.Router.Group("/api/auth")
	{
		api.POST("/signup", u.AuthController.SignUp)
		api.POST("/signin", u.AuthController.SignIn)
	}

}

func NewAuthRoute(authController controllers.IAuthController, router *gin.Engine) *AuthRoute {
	return &AuthRoute{AuthController: authController, Router: router}
}
