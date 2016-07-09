package main

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/ankalus/contrib/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", func(c *gin.Context) {
		// Create the token with some claims
		token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, jwt_lib.MapClaims{
			"ID": "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		})
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(mysupersecretpassword))
		if err != nil {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
	})

	private := r.Group("/api/private")
	private.Use(jwt.Auth(mysupersecretpassword))

	/*
		Set this header in your request to get here.
		Authorization: Bearer `token`
	*/

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})

	r.Run("localhost:8080")
}
