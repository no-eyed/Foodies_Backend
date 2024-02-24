package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"foodiesbackend/utils"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		err := godotenv.Load()

		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		client_secret_key := os.Getenv("CLERK_SECRET_KEY")

		client, err := clerk.NewClient(client_secret_key)

		if err != nil {

			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Error creating Clerk client",
			})

			return

		}

		token := strings.TrimPrefix(context.GetHeader("Authorization"), "Bearer ")

		sessionId, err := utils.GetSessionIDFromToken(token)
		if err != nil {
			fmt.Println("Error getting session ID from token:", err)

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not authorized",
			})

			return

		}

		session, err := client.Sessions().Read(sessionId)

		if err != nil || session.Status != "active" {

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired session",
			})

			return
		}

		context.Next()
	}
}

func VerifyUser() gin.HandlerFunc {
	return func(context *gin.Context) {

		err := godotenv.Load()

		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		client_secret_key := os.Getenv("CLERK_SECRET_KEY")

		client, err := clerk.NewClient(client_secret_key)

		if err != nil {

			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Error creating Clerk client",
			})

			return

		}

		token := strings.TrimPrefix(context.GetHeader("Authorization"), "Bearer ")

		sessionId, err := utils.GetSessionIDFromToken(token)

		if err != nil {
			fmt.Println("Error getting session ID from token:", err)

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not authorized",
			})

			return

		}

		session, err := client.Sessions().Read(sessionId)

		clerkid := context.Param("clerkid")

		if err != nil || session.Status != "active" || session.UserID != clerkid {

			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired session",
			})

			return
		}

		context.Next()
	}
}
