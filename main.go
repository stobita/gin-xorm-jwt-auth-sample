package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stobita/gin-xorm-jwt-auth-sample/handler"
	"github.com/stobita/gin-xorm-jwt-auth-sample/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/", handler.AnonymousHello())
	r.POST("/signup", handler.UserSignUp())
	r.POST("/signin", handler.UserSignIn())
	authorized := r.Group("/", middleware.TokenAuthMiddleware())
	{
		authorized.GET("/private", handler.PrivateHello())
	}
	// for live reload
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
