package middlewares

import (
	jwtUtil "chilindo/pkg/utils"
	"chilindo/src/user-service/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type IJwtMiddleWare interface {
	MiddleWare() gin.HandlerFunc
}

type JwtMiddleWare struct {
	tokenController *jwtUtil.JWTClaim
}

func (s *JwtMiddleWare) IsAuth() gin.HandlerFunc {
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

		claims, err := jwtUtil.ExtractToken(tokenResult)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token is expired",
			})
			log.Println("Error: Token is expired")
			c.Abort()
			return
		}
		c.Set(config.UserID, claims.Id)
		c.Next()
	}
}
