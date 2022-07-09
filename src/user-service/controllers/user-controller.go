package controllers

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/services"
	"chilindo/src/user-service/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const addressId = "addressId"

type IUserController interface {
	GetAddress(c *gin.Context)
	GetAddressById(c *gin.Context)
	CreateAddressByUserId(c *gin.Context)
	DeleteAddressById(c *gin.Context)
}

type UserController struct {
	UserService services.IUserService
	Token       *token.JWTClaim
}

func (u *UserController) CreateAddressByUserId(c *gin.Context) {
	var dTo *dto.CreateAddressDTO
	userId, oke := c.Get(config.UserID)
	if !oke {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddressByUserId: Error Get User ID in package controller")
		c.Abort()
		return
	}
	dTo.UserId = userId.(int)
	address, err := u.UserService.CreateAddress(dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error Add address",
		})
		log.Println("CreateAddressByUserId: Error create new address in package controller")
		return
	}
	c.JSONP(http.StatusOK, address)
}

func (u *UserController) GetAddress(c *gin.Context) {
	var dTo *dto.GetAddressDTO
	userId, oke := c.Get(config.UserID)
	dTo.UserId = userId.(int)
	if !oke {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address is fail",
		})
		log.Println("GetAddress: Error Get Address in package controller")
		c.Abort()
		return
	}
	address, err := u.UserService.GetAddress(dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address is fail",
		})
		log.Println("GetAddress: Error Get Address in package controller")
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, address)
}

func (u *UserController) GetAddressById(c *gin.Context) {
	var dTo *dto.GetAddressByIdDTO
	addressId, errCv := strconv.Atoi(c.Param(addressId))
	if errCv != nil {
		c.JSONP(http.StatusBadRequest, gin.H{})
		log.Println("GetAddressById: Can't get addressId")
		return
	}
	userId, ok := c.Get(config.UserID)
	if !ok {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address by userId fail",
		})
		log.Println("GetAddressById: Error to get Id form user in package controllers")
		return
	}
	dTo.AddressId = addressId
	dTo.UserId = userId.(int)
	address, err := u.UserService.GetAddressById(dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address by ID fail",
		})
		log.Println("GetAddressById: Error in package controllers", err)
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, address)
}

func (u *UserController) DeleteAddressById(c *gin.Context) {
	var dTo *dto.DeleteAddressByIdDTO
	param := c.Param(addressId)
	addressID, errGetAddressId := strconv.Atoi(param)
	if errGetAddressId != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Fail to Delete Address",
		})
		log.Println("DeleteAddressById: Error to get addressID in package controller", errGetAddressId)
		c.Abort()
		return
	}
	userId, ok := c.Get(config.UserID)
	if !ok {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Fail to Delete Address",
		})
		log.Println("DeleteAddressById: Error to get userId in package controller")
		c.Abort()
		return
	}
	dTo.UserId = userId.(int)
	dTo.AddressId = addressID
	address, err := u.UserService.DeletedAddressById(dTo)
	if err != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Fail to Delete Address",
		})
		log.Println("DeleteAddressById: Error to delete Address in package controller")
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, address)
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{UserService: userService}
}
