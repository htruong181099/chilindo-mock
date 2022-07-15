package client

import (
	pb "chilindo/pkg/pb/admin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckIsAdmin(adminClient pb.AdminServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			c.Abort()
			return
		}
		res, err := adminClient.CheckIsAdmin(c, &pb.CheckIsAdminRequest{
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
		log.Println("isAuth: ", res.IsAuth)
		log.Println("isAdmin: ", res.IsAdmin)
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
