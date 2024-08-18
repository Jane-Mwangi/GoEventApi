	package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Jane-Mwangi/GoEventApi/models"
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
	log.Println("Reached createEvent route handler")

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	log.Println("Creating event with UserID:", userId)

	event.UserID = userId // Set the user ID to the authenticated user``
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return

	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully", " event": event})
}

func updateEvent(context *gin.Context) {
	log.Println("Reached updateEvent route handler")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
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
	log.Println("Reached deleteEvent route handler")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event ID", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	log.Println("Extracted userId from context:", userId)

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.Try Again Later", "error": err.Error()})
		return
	}
	log.Println("Event retrieved:", event)
	log.Println("UserID associated with event:", event.UserID)

	if event.UserID != userId {

		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event.Try Again Later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully", "event": event})
}
