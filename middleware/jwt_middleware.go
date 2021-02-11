package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	tokenSecret = []byte("dwfwe")
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("x-auth-token")

		if tokenString == "" {

			c.Abort()

		} else {

			result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return tokenSecret, nil
			})

			if err == nil && result.Valid {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"auth": false,
				})
				c.Abort()
			}

		}

	}

}
