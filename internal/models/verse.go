package models

import "time"

// Verse represents a single verse from the srimad bhagavatam
type Verse struct {
	ID            int       `json:"id"`
	CantoNumber   int       `json:"canto_number"`
	ChapterNumber int       `json:"chapter_number"`
	VerseNumber   int       `json:"verse_number"`
	Translation   string    `json:"translation"`
	Purport       string    `json:"purport"`
	CreatedAt     time.Time `json:"created_at"`
}
