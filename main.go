package main

import (
	"foodiesbackend/db"
	"foodiesbackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Initdb()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
