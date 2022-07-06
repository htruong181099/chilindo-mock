package controllers

import (
	"chilindo/dto"
	"chilindo/models"
	"chilindo/services"
	"chilindo/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAuthController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}

type AuthController struct {
	AuthService services.IAuthService
	Token       *token.JWTClaim
}

func (u AuthController) SignIn(c *gin.Context) {
	var user *dto.SignInDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error to sign in",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", err)
		return
	}
	userLogin, errLogin := u.AuthService.SignIn(user)
	if errLogin != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "Error SignIn",
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		return
	}
	tokenString, errToken := u.Token.GenerateJWT(userLogin.Username, userLogin.Email, userLogin.Id)
	if errToken != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error SignIn",
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}

func (u AuthController) SignUp(c *gin.Context) {
	var userBody *models.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&userBody); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error to sign up",
		})
		log.Println("SignUp: Error ShouldBindJSON in package controller", err.Error())
		return
	}
	user, err := u.AuthService.SignUp(userBody)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSONP(http.StatusOK, user)
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{AuthService: authService}
}
