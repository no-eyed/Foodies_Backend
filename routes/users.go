package routes

import (
	"foodiesbackend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser(context *gin.Context) {
	clerkId := context.Param("clerkid")
	userId, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id"})
		return
	}

	user, err := models.GetUserById(userId)

	if err == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func ManageUser(context *gin.Context) {
	var user models.User

	var reqUser models.UserRequest

	err := context.BindJSON(&reqUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request", "error": err.Error()})
		return
	}
	user.ClerkId = reqUser.Data.UUID
	user.Id, err = models.GetUserIdByClerkid(user.ClerkId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user", "error": err.Error()})
		return
	}

	if reqUser.Type == "user.deleted" {
		err = user.Delete()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete user", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, user)
		return
	}
	user.Username = reqUser.Data.UserName
	user.Email = reqUser.Data.Email[0].EmailAddress
	user.CreatedAt = time.Unix(reqUser.Data.CreatedAt/1000, 0)
	user.UpdatedAt = time.Unix(reqUser.Data.UpdatedAt/1000, 0)

	if reqUser.Type == "user.updated" {
		err = user.Update()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, user)
		return
	} else if reqUser.Type == "user.created" {
		err = user.Save()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, user)
		return
	}
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
