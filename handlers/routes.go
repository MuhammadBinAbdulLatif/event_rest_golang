package handlers

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/events", CreateEvent)
	server.GET("/events/:id", GetEvent)
	authenticatedRoutes.PUT("/events/:id", UpdateEvent)
	authenticatedRoutes.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", SignupUser)
	server.POST("/login", LoginUser)
	authenticatedRoutes.POST("/events/:id/register", RegisterForEvent)
	authenticatedRoutes.DELETE("/events/:id/register", UnregisterForEvent)	
}

