package services

import (
	"fmt"
	"net/http"
	"proyectos/servidor-cliente/authorization"
	"proyectos/servidor-cliente/database"
	"proyectos/servidor-cliente/models"

	"github.com/gin-gonic/gin"
)

var (
	secret = []byte("secret")
)

func Login(c *gin.Context) {

	database.DBConnection()

	var reqBody models.Usuario

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request Body",
		})
		return
	}

	err := database.DBClient.Get(&reqBody, "SELECT correo,password FROM usuarios WHERE correo = ? AND password = ?",
		reqBody.Correo,
		reqBody.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"auth":     false,
			"email":    reqBody.Correo,
			"password": reqBody.Password,
			"error":    err.Error(),
		})
		return
	}

	token, err := authorization.GetToken(reqBody.Correo, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
