package routes

import (
	"foodiesbackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	clerkId := context.Param("clerkid")
	userId, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id"})
	}

	user, err := models.GetUserById(userId)

	if err == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
	}

	context.JSON(http.StatusOK, user)
}

func CreateUser(context *gin.Context) {
	user := models.User{}
	context.BindJSON(&user)

	err := user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	clerkId := context.Param("clerkid")
	userId, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}

	context.BindJSON(&user)

	err = user.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func DeleteUser(context *gin.Context) {
	clerkId := context.Param("clerkid")
	userId, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id"})
		return
	}

	user, err := models.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}

	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}
