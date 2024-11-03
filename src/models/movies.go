package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Movies struct {
	ID          uint      `gorm:"primary_key"`
	UniqueID    uuid.UUID `gorm:"unique"`
	Title       string    `gorm:"size:50;not null"`
	Description string
	Duration    int64          `gorm:"not null"`
	Artists     datatypes.JSON `gorm:"type:json"`
	GenreIds    datatypes.JSON `gorm:"type:json"`
	WatchUrl    string         `gorm:"not null"`
	Viewers     int            `gorm:"not null"`
	Voters      int            `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
