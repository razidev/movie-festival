package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uint      `gorm:"primary_key"`
	UniqueID  uuid.UUID `gorm:"unique"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"password;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
