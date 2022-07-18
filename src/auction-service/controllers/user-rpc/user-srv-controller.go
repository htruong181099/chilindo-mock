package controllers

import (
	"chilindo/pkg/configs"
	adminPb "chilindo/pkg/pb/admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserAuthServiceController interface {
	CheckIsAuth() gin.HandlerFunc
}

type UserAuthServiceController struct {
	UserServiceClient adminPb.AdminServiceClient
}

func NewUserAuthServiceController(userServiceClient adminPb.AdminServiceClient) *UserAuthServiceController {
	return &UserAuthServiceController{UserServiceClient: userServiceClient}
}

func (a UserAuthServiceController) CheckIsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			c.Abort()
			return
		}
		fmt.Println("Check err", a.UserServiceClient)

		res, err := a.UserServiceClient.CheckUserAuth(c, &adminPb.CheckUserAuthRequest{
			Token: token,
		})
		//fmt.Println("Check err", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error call admin service",
			})
			c.Abort()
			return
		}
		userId := int(res.UserId)
		c.Set(configs.UserID, userId)
		if !res.IsAuth {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
