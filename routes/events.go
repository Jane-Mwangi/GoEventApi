package routes

import (
	"net/http"
	"strconv"

	"github.com/Jane-Mwangi/GoEventApi/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.Try Again Later", "error": err.Error()})
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
	}
	context.JSON(http.StatusOK, event)

}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.Try Again Later", "error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully", " event": event})
}
