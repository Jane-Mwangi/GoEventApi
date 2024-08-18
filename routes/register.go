package routes

import (
	"net/http"
	"strconv"

	"github.com/Jane-Mwangi/GoEventApi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})


}

func unregisterForEvent(context *gin.Context) {

}
