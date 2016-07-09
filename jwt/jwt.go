package jwt

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	jwt_request_lib "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token
		_, err := jwt_request_lib.ParseFromRequest(c.Request, ArgumentExtractor{"token"}, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}
