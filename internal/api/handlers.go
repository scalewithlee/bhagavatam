package api

import (
	"bhagavatam/internal/database"
	"bhagavatam/internal/models"
	"log"
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

// CreateVerseHandler creates a new verse in the database
func CreateVerseHandler(c *gin.Context) {
	var newVerse models.Verse

	// Parse JSON request body into our struct
	if err := c.ShouldBindJSON(&newVerse); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse JSON request body"})
		return
	}

	// Basic validation
	if newVerse.CantoNumber < 1 || newVerse.CantoNumber > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "canto must be between 1 and 12"})
		return
	}
	if newVerse.ChapterNumber < 1 || newVerse.ChapterNumber > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chapter number"})
		return
	}
	if newVerse.VerseNumber < 1 || newVerse.VerseNumber > 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verse number"})
		return
	}
	if newVerse.Translation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "translation cannot be empty"})
	}

	// Insert into database
	result := database.DB.Create(&newVerse)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create verse"})
		return
	}

	// Return the created verse with ID populated
	c.JSON(http.StatusCreated, newVerse)
}
