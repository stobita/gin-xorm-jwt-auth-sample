package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stobita/gin-xorm-jwt-auth-sample/lib"
	"github.com/stobita/gin-xorm-jwt-auth-sample/model"
)

type SignInJSON struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpJSON struct {
	Name                 string `json:"name" binding:"required"`
	Password             string `json:"password" binding:"required"`
	ConfirmationPassword string `json:"confirmationPassword" binding:"required"`
}

func UserSignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json SignInJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, gin.H{"message": "Invalid Params"})
			return
		}
		name := json.Name
		password := json.Password
		result := model.User{Name: name}.GetUser()
		if result == nil {
			c.JSON(400, gin.H{"message": "User not found"})
			return
		}
		log.Println(result)
		if lib.ComparePassword(name, password, result.Password) {
			if tokenString, err := lib.GetTokenString(name); err == nil {
				c.JSON(200, gin.H{"token": tokenString})
				return
			}
		}
		c.JSON(400, gin.H{"message": "Invalid name or password"})
	}
}

func UserSignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json SignUpJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, gin.H{"message": "Invalid Params"})
			return
		}
		name := json.Name
		password := json.Password
		confirmationPassword := json.ConfirmationPassword
		if password != confirmationPassword {
			c.JSON(400, gin.H{"message": "Password mismatch"})
			return
		}
		encryptedPassword, err := lib.GetEncryptedPassword(password)
		if err != nil || encryptedPassword == "" {
			c.JSON(400, gin.H{"message": "Could not encrypt password"})
			return
		}
		result := model.User{Name: name, Password: encryptedPassword}.Insert()
		c.JSON(200, gin.H{"result": result != nil})
	}
}
