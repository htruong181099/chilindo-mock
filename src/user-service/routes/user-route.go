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

	}

	addressAPI := u.Router.Group("/api/users/address").Use(u.JWTMiddleware.IsAuth())
	{
		addressAPI.GET("/", u.UserController.GetAddress)
		addressAPI.GET("/:addressId", u.UserController.GetAddressById)
		addressAPI.POST("/", u.UserController.CreateAddressByUserId)
		addressAPI.PATCH("/:addressId", u.UserController.UpdateAddressById)
		addressAPI.DELETE("/:addressId", u.UserController.DeleteAddressById)
	}

}

func NewUserRoute(userController controllers.IUserController, router *gin.Engine) *UserRoute {
	return &UserRoute{UserController: userController, Router: router}
}
