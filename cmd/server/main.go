package main

import (
	"bhagavatam/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router with default middleware (logger and recovery)
	// router is a *gin.Engine, which is like a manager
	router := gin.Default()

	// Endpoints (like employees, who will return back to the manager, and the manager returns the response)
	router.GET("/health", api.HealthHander)
	router.GET("/verses/:canto/:chapter/:verse", api.GetVerseHandler)

	// Start up the server
	router.Run(":8080")
}
