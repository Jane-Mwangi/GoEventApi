package routes

import (

	"net/http"
	"strconv"

	"github.com/Jane-Mwangi/GoEventApi/models"
	"github.com/Jane-Mwangi/GoEventApi/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
		return
	}

	err:=utils.VerifyToken(token)

	if err!=nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully", " event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
		return
	}

	_, err = models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully", "event": updateEvent})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully", "event": event})
}
