package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/middlewares"
)

func Register(server *gin.Engine) {
	server.GET("/events", getEvents)

	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	{
		protected.POST("/events", createEvent)
		protected.GET("/events/:id", getEvent)
		protected.PUT("/events/:id", updateEvent)
		protected.DELETE("/events/:id", deleteEvent)

		protected.POST("/events/:id/registration", registerForEvent)
		protected.DELETE("/events/:id/registration", unregisterFromEvent)
	}

	server.POST("/signup", signup)
	server.POST("/login", login)
}
