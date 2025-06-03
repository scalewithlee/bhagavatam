package api

import (
	"bhagavatam/internal/database"
	"bhagavatam/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HealthHandler returns the health of the api
func HealthHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "bhagavatam-api",
	})
}

// GetVerseHandler returns a specific verse (hard-coded for now)
func GetVerseHandler(c *gin.Context) {
	// Extract the URL parameters
	cantoStr := c.Param("canto")
	chapterStr := c.Param("chapter")
	verseStr := c.Param("verse")

	// Convert strings to integers
	canto, err := strconv.Atoi(cantoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid canto number"})
		return
	}
	chapter, err := strconv.Atoi(chapterStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chapter number"})
		return
	}
	verse, err := strconv.Atoi(verseStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verse number"})
		return
	}

	// There are only 12 cantos in the bhagavatam
	if canto < 1 || canto > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "canto must be between 1 and 12"})
		return
	}

	if chapter < 1 || chapter > 100 { // rough upper limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chapter number"})
		return
	}

	if verse < 1 || verse > 200 { // rough upper limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verse number"})
		return
	}

	// Query the database
	var foundVerse models.Verse
	result := database.DB.Where("canto_number = ? AND chapter_number = ? AND verse_number = ?", canto, chapter, verse).First(&foundVerse)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "verse not found"})
		return
	}

	c.JSON(http.StatusOK, foundVerse)
}
