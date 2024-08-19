package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Jane-Mwangi/GoEventApi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
		return
	}

	token = strings.TrimSpace(token)

	// // Clean the token by removing any extraneous characters or formatting
	// token = strings.TrimSpace(token)
	// if strings.HasPrefix(token, "Bearer ") {
	// 	token = strings.TrimPrefix(token, "Bearer ")
	// }

	
	// log.Println("Cleaned Authorization token:", token)

	// Log the token to ensure it's being retrieved correctly
	log.Println("Authorization token:", token)

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}