package authorization

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	tokenSecret = []byte("dwfwe")
)

func GetToken(correo string, c *gin.Context) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": correo,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	strToken, err := token.SignedString(tokenSecret)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	return strToken, nil
}
