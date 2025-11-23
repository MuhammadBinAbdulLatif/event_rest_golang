package handlers

import (
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(c *gin.Context) {
	eventId, err:= strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventByID(strconv.FormatInt(eventId, 10))
	if err != nil {
		c.JSON(500, gin.H{"error": "Event not found"})
		return
	}
	userId := c.GetInt64("userId")
	 err=event.Register(userId)
	 if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Registered for event successfully"})
}


func UnregisterForEvent(c *gin.Context) {
		eventId, err:= strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var event models.Event
	event.ID = int(eventId)
	userId := c.GetInt64("userId")
	 err=event.Unregister(int64(userId))
	 if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Unregistered for event successfully"})
}