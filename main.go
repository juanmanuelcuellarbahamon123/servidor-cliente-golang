package main

import (
	"proyectos/servidor-cliente/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	routes.UsersRoutes(router)

	router.Run(":8000")

}
