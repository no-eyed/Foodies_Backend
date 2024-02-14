package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/meals", GetMeals)
	server.GET("/meals/:id", GetMeal)
	server.POST("/meals", CreateMeal)
	server.PUT("/meals/:id", UpdateMeal)
	server.DELETE("/meals/:id", DeleteMeal)

	// server.GET("/events/:id", getEvent)

	// server.POST("/signup", signup)
	// server.POST("/login", login)
}
