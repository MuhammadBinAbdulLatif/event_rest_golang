package handlers

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)



func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, events)
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	event.UserId = userId
	err := event.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return	
	}
	c.JSON(201, gin.H{"message": "Event created successfully"})
}


func GetEvent(c *gin.Context) {
	id := c.Param("id")
	event,err:= models.GetEventByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"event": event})
}


func UpdateEvent(c *gin.Context) {
	
	eventId, err:= strconv.ParseInt(c.Param("id"), 10, 64)
	event_id := c.Param("id")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	var event models.Event;
	event, err = models.GetEventByID(event_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var updatedEvent models.Event
	err=c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	if userId != event.UserId {
		c.JSON(401, gin.H{"error": "You are unauthorized to update this event"})
		return
	}
	updatedEvent.ID =int(eventId)
	// let's update the event
	err=updatedEvent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated succesfully"})

}


func DeleteEvent(c *gin.Context) {
	_, err:= strconv.ParseInt(c.Param("id"), 10, 64)
	event_id := c.Param("id")
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot parse the id param"})
	}
	event, err := models.GetEventByID(event_id)
	userId := c.GetInt64("userId")
	if userId != event.UserId {
		c.JSON(401, gin.H{"error": "You are unauthorized to delete this event"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot find the event to delete"})
		return
	}
	err =event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete events"})

	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})


}