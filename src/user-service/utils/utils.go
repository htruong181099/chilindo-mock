package utils

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type IMiddleWare interface {
	MiddleWare() gin.HandlerFunc
}

type SMiddleWare struct {
	tokenController *token.JWTClaim
}

func (s *SMiddleWare) MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSONP(http.StatusUnauthorized, gin.H{
				"Message": "Request doest not contain token",
			})
			log.Println("MiddleWare: Error to get token in")
			c.Abort()
			return
		}
		tokenResult := strings.TrimPrefix(tokenString, "Bearer ")
		fmt.Println("Check Token", tokenResult)
		claims := token.ExtractToken(tokenResult)
		if claims.ExpiresAt < time.Now().Local().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token is expired",
			})
			log.Println("Error: Token is expired")
			c.Abort()
			return
		}
		c.Set(config.UserID, claims.Id)
		//claims := s.tokenController.ExtractToken(tokenResult)
		fmt.Println(claims)
		c.Next()
	}
}
