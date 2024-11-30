package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oyen-bright/event_REST/db"
	"github.com/oyen-bright/event_REST/routes"
)

func main() {
	db.Init()
	server := gin.Default()

	routes.Register(server)

	// server.GET("/events/html", getEventHTML)
	server.Run(":8080")
}

// func getEventHTML(context *gin.Context) {
// 	events := models.GetAllEvents()
// 	html := "<ul>"
// 	for _, event := range events {
// 		html += "<li>" + event.Name + "</li>"
// 	}
// 	html += "</ul>"
// 	context.Header("Content-Type", "text/html")
// 	context.String(http.StatusOK, html)
// }
