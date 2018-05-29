package lib

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(name string, receivedPassword string, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(receivedPassword))
	return err == nil
}

func GetTokenString(name string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"sub": name,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	if tokenString, err := token.SignedString([]byte(secretKey)); err == nil {
		return tokenString, nil
	} else {
		return "", err
	}
}

func GetEncryptedPassword(rawPassword string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err == nil {
		return string(encryptedPassword), nil
	} else {
		return "", nil
	}
}

func TokenAuthenticate(tokenString string) bool {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if tokenString == "" {
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["sub"] == nil {
			return false
		}
	} else {
		return false
	}
	return true
}
