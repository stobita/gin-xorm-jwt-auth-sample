package handler

import "github.com/gin-gonic/gin"

func AnonymousHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Anonymous User"})
	}
}
func PrivateHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Private User"})
	}
}
