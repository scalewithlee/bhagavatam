package main

import (
	"bhagavatam/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Endpoints
	router.GET("/health", api.HealthHander)
	router.GET("/verses/:canto/:chapter/:verse", api.GetVerseHandler)

	// Start up the server
	router.Run(":8080")
}
