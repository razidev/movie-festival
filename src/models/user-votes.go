package models

import (
	"time"

	"github.com/google/uuid"
)

type UserVotes struct {
	ID            uint      `gorm:"primary_key"`
	MovieUniqueID uuid.UUID `gorm:"not null"`
	UserUniqueID  uuid.UUID `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
