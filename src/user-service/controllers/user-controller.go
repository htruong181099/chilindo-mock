package controllers

import (
	"chilindo/services"
	"chilindo/token"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetAddress(c *gin.Context)
}

type UserController struct {
	UserService services.IUserService
	Token       *token.JWTClaim
}

func (u UserController) GetAddress(c *gin.Context) {
	//TODO implement me
	//addressId := c.Param("addressId")
	//userId, ok := c.Get("userId")
	//var dto = &dto.GetAddressDTO{addressId: addressId,userId: userId}
	//u.UserService.GetAddress(dto)
	panic("implement me")
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{UserService: userService}
}
