package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stobita/gin-xorm-jwt-auth-sample/lib"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(400, gin.H{"message": "Auth header empty"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if !lib.TokenAuthenticate(tokenString) {
			c.JSON(400, gin.H{"message": "Invalid Token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
