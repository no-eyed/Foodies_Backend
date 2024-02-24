package routes

import (
	"foodiesbackend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMeals(context *gin.Context) {
	//meal := models.Meal{}
	meals, err := models.GetAllMeals()

	// fmt.Println("i am the architect of my own destruction")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meals", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meals)
}

func GetMeal(context *gin.Context) {
	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse meal id"})
		return
	}

	meal, err := models.GetResponseMealById(mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func CreateMeal(context *gin.Context) {
	meal := models.Meal{}
	context.BindJSON(&meal)

	clerkId := context.Param("clerkid")
	Id, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse uid", "error": err.Error()})
		return
	}

	meal.Creator_id = Id

	err = meal.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func UpdateMeal(context *gin.Context) {
	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse meal id"})
		return
	}

	meal, err := models.GetMealById(mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meal", "error": err.Error()})
		return
	}

	context.BindJSON(&meal)

	err = meal.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func DeleteMeal(context *gin.Context) {
	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse meal id"})
		return
	}

	meal, err := models.GetMealById(mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meal", "error": err.Error()})
		return
	}

	err = meal.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete meal", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})
}

func GetMyMeals(context *gin.Context) {
	clerkId := context.Param("clerkid")
	userID, err := models.GetUserIdByClerkid(clerkId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meals", "error": err.Error()})
	}

	meals, err := models.GetMealsByCreatorId(userID)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meals", "error": err.Error()})
	}

	context.JSON(http.StatusOK, meals)
}
