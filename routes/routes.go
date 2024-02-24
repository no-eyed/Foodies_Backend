package routes

import (
	middlewares "foodiesbackend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/meals", GetMeals)
	server.GET("/meals/:id", GetMeal)
	server.POST("/meals/:clerkid", middlewares.AuthMiddleware(), CreateMeal)
	server.PUT("/meals/:clerkid/:id", middlewares.VerifyUser(), UpdateMeal)
	server.DELETE("/meals/:clerkid/:id", middlewares.VerifyUser(), DeleteMeal)
	server.GET("/meals/my-meals/:clerkid", middlewares.VerifyUser(), GetMyMeals)

	server.GET("/user/:clerkid", GetUser)
	server.POST("/user", CreateUser)
	server.PUT("/user/:clerkid", UpdateUser)
	server.DELETE("/user/:clerkid", DeleteUser)
}
