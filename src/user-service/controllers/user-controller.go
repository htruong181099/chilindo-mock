package controllers

import (
	jwtUtil "chilindo/pkg/utils"
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/models"
	"chilindo/src/user-service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	addressId = "addressId"
)

type IUserController interface {
	GetUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	GetAddress(c *gin.Context)
	GetAddressById(c *gin.Context)
	CreateAddressByUserId(c *gin.Context)
	DeleteAddressById(c *gin.Context)
	UpdateAddressById(c *gin.Context)
}

type UserController struct {
	UserService services.IUserService
	Token       *jwtUtil.JWTClaim
}

func (u *UserController) ChangePassword(c *gin.Context) {
	var dTo dto.UpdatePasswordDTO
	userId, ok := c.Get(config.UserID)
	if !ok {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Unauthorized",
		})
		log.Println("ChangePassword: Error Get UserID in package controller")
		c.Abort()
		return
	}
	if err := c.ShouldBindJSON(&dTo); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error to change password",
		})
		log.Println("ChangePassword: Error ShouldBindJSON in package controller")
		c.Abort()
		return
	}
	if len(dTo.NewPassword) < 6 || dTo.CurrentPassword == dTo.NewPassword {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Invalid Password",
		})
		c.Abort()
		return
	}
	//fmt.Println("Check here")

	dTo.UserId = userId.(int)
	_, err := u.UserService.UpdatePassword(&dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error update password",
		})
		log.Println("ChangePassword: Error ChangePassword in package controller", err)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "Success to change your password",
	})
}

func (u *UserController) GetUser(c *gin.Context) {
	var dTo dto.GetByUserIdDTO
	userId, ok := c.Get(config.UserID)
	fmt.Println(userId)
	if !ok {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Unauthorized",
		})
		log.Println("CreateAddressByUserId: Error Get UserID in package controller")
		c.Abort()
		return
	}
	dTo.UserId = userId.(int)
	user, err := u.UserService.GetUserById(&dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error get user",
		})
		log.Println("GetUser: Error get user in package controller", err)
		return
	}
	if user == nil {
		c.JSONP(http.StatusNotFound, gin.H{
			"Message": "Not found",
		})
		return
	}
	user.Password = ""
	c.JSONP(http.StatusOK, gin.H{
		"id":          user.Id,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"username":    user.Username,
		"email":       user.Email,
		"phoneNumber": user.PhoneNumber,
		"gender":      user.Gender,
	})
}

func (u *UserController) CreateAddressByUserId(c *gin.Context) {
	//var dTo dto.AddressDTO
	userId, oke := c.Get(config.UserID)
	if !oke {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddressByUserId: Error Get User ID in package controller")
		c.Abort()
		return
	}
	var addressBody *models.Address
	if err := c.ShouldBindJSON(&addressBody); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddressByUserId: Error ShouldBindJSON in package controller")
		c.Abort()
		return
	}
	dTo := dto.NewAddressDTO(addressBody)
	dTo.Address.UserId = userId.(int)
	address, err := u.UserService.CreateAddress(dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error Add address",
		})
		log.Println("CreateAddressByUserId: Error create new address in package controller")
		return
	}
	c.JSONP(http.StatusCreated, address)
}

func (u *UserController) GetAddress(c *gin.Context) {
	var dTo dto.GetAddressDTO
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
	address, err := u.UserService.GetAddress(&dTo)
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
	var dTo dto.GetAddressByIdDTO
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
	address, err := u.UserService.GetAddressById(&dTo)
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
	var dTo dto.GetAddressByIdDTO
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
	address, err := u.UserService.DeletedAddressById(&dTo)
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

func (u *UserController) UpdateAddressById(c *gin.Context) {
	userId, oke := c.Get(config.UserID)
	if !oke {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Error Update Address",
		})
		log.Println("CreateAddressByUserId: Error Get User ID in package controller")
		c.Abort()
		return
	}
	var addressBody *models.Address
	if err := c.ShouldBindJSON(&addressBody); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error update address",
		})
		log.Println("CreateAddressByUserId: Error ShouldBindJSON in package controller")
		c.Abort()
		return
	}

	addressId, errCv := strconv.Atoi(c.Param(addressId))
	if errCv != nil {
		c.JSONP(http.StatusBadRequest, gin.H{})
		log.Println("UpdateAddressById: Can't get addressId")
		return
	}

	dTo := dto.NewAddressDTO(addressBody)
	dTo.Address.UserId = userId.(int)
	dTo.Address.Id = addressId
	address, err := u.UserService.UpdateAddressById(dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error Add address",
		})
		log.Println("CreateAddressByUserId: Error create new address in package controller")
		return
	}
	c.JSONP(http.StatusOK, address)
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{UserService: userService}
}
