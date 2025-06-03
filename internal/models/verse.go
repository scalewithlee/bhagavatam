package models

import (
	"time"

	"gorm.io/gorm"
)

// Verse represents a single verse from the srimad bhagavatam
type Verse struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	CantoNumber   int            `json:"canto_number" gorm:"column:canto_number;not null"`
	ChapterNumber int            `json:"chapter_number" gorm:"column:chapter_number;not null"`
	VerseNumber   int            `json:"verse_number" gorm:"column:verse_number;not null"`
	Translation   string         `json:"translation" gorm:"type:text;not null"`
	Purport       string         `json:"purport,omitempty" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
