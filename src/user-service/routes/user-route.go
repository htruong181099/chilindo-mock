package routes

import (
	"chilindo/pkg/middlewares"
	"chilindo/src/user-service/controllers"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	SetRouter()
}
type UserRoute struct {
	UserController controllers.IUserController
	Router         *gin.Engine
	JWTMiddleware  *middlewares.JwtMiddleWare
}

func (u UserRoute) SetRouter() {

	api := u.Router.Group("/api/users").Use(u.JWTMiddleware.IsAuth())
	{
		api.GET("/", u.UserController.GetUser)
		api.PATCH("/password", u.UserController.ChangePassword)
		api.GET("/address", u.UserController.GetAddress)
		api.GET("/address/:addressId", u.UserController.GetAddressById)
		api.POST("/address", u.UserController.CreateAddressByUserId)
		api.DELETE("/address/:addressId", u.UserController.DeleteAddressById)
		api.PATCH("/address/:addressId", u.UserController.UpdateAddressById)
	}

}

func NewUserRoute(userController controllers.IUserController, router *gin.Engine) *UserRoute {
	return &UserRoute{UserController: userController, Router: router}
}
