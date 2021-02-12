package services

import (
	"fmt"
	"net/http"
	"proyectos/servidor-cliente/database"
	"proyectos/servidor-cliente/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListarUsuarios(c *gin.Context) {

	database.DBConnection()

	var usuarios []models.Usuario

	err := database.DBClient.Select(&usuarios, "SELECT nombre,apellido,correo,password,id_rol FROM usuarios")

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, usuarios)

}

func ListarUsuario(c *gin.Context) {

	database.DBConnection()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	var usuario models.Usuario

	err2 := database.DBClient.Get(&usuario, "SELECT nombre,apellido,correo,password,id_rol FROM usuarios WHERE id_usuario = ?", id)

	if err2 != nil {
		c.JSON(401, gin.H{
			"msg": "no se encontro resultado",
		})
	}

	c.JSON(200, usuario)

}

func EliminarUsuario(c *gin.Context) {

	database.DBConnection()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	res, err2 := database.DBClient.Exec("DELETE FROM usuarios WHERE id_usuario = ?", id)

	if err2 != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"msg": "El usuario ha sido eliminado",
		"res": res,
	})

}

func ActualizarUsuario(c *gin.Context) {

	database.DBConnection()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	var reqBody models.Usuario

	if err2 := c.ShouldBindJSON(&reqBody); err2 != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err2.Error(),
			"msg":   "json no valido",
		})
	}

	res, err3 := database.DBClient.Exec("UPDATE usuarios SET nombre = ?, apellido = ?, correo = ?, password = ? WHERE id_usuario = ?",
		reqBody.Nombre,
		reqBody.Apellido,
		reqBody.Correo,
		reqBody.Password,
		id,
	)

	if err3 != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"msg": "El usuario se ha modificado",
		"res": res,
	})

}

func AgregarUsuario(c *gin.Context) {

	database.DBConnection()

	var reqBody models.Usuario

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
			"msg":   "json no valido",
		})
	}

	var email int
	err2 := database.DBClient.Get(&email, "SELECT COUNT(*) FROM usuarios WHERE correo = ?", reqBody.Correo)

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	if email > 0 {
		c.JSON(401, gin.H{
			"msg":   "Este correo ya existe",
			"email": email,
		})
		return
	}

	_, err3 := database.DBClient.Exec("INSERT INTO usuarios(nombre,apellido,correo,password,id_rol) VALUES (?,?,?,?,?)",
		reqBody.Nombre,
		reqBody.Apellido,
		reqBody.Correo,
		reqBody.Password,
		reqBody.Rol,
	)

	if err3 != nil {
		fmt.Println(err3.Error())
	}

	c.JSON(200, gin.H{
		"msg": "usuario agregado",
	})

}
