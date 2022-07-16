package controllers

import (
	adminPb "chilindo/pkg/pb/admin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAdminServiceController interface {
	CheckIsAdmin() gin.HandlerFunc
}

type AdminServiceController struct {
	AdminClient adminPb.AdminServiceClient
}

func NewAdminServiceController(adminClient adminPb.AdminServiceClient) *AdminServiceController {
	return &AdminServiceController{AdminClient: adminClient}
}

func (a AdminServiceController) CheckIsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			c.Abort()
			return
		}
		res, err := a.AdminClient.CheckIsAdmin(c, &adminPb.CheckIsAdminRequest{
			Token: token,
		})
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Error call admin service",
			})
			c.Abort()
			return
		}
		if !(res.IsAuth && res.IsAdmin) {
			c.JSON(http.StatusForbidden, gin.H{
				"Message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
