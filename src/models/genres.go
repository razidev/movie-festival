package models

import "time"

type Genres struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Viewers   int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
