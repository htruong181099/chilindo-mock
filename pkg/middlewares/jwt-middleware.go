package middlewares

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/token"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type IJwtMiddleWare interface {
	MiddleWare() gin.HandlerFunc
}

type JwtMiddleWare struct {
	tokenController *token.JWTClaim
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
		//fmt.Println("Check Token", tokenResult)

		claims, err := token.ExtractToken(tokenResult)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token is expired",
			})
			log.Println("Error: Token is expired")
			c.Abort()
			return
		}
		c.Set(config.UserID, claims.Id)
		//claims := s.tokenController.ExtractToken(tokenResult)
		//fmt.Println(claims)
		c.Next()
	}
}
