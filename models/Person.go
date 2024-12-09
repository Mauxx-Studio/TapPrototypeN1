package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	RealID     string    `gorm:"not null" json:"real_id"`
	InternalID string    `json:"internal_id"`
	Birthdate  string    `json:"birthdate"`
	Gender     string    `json:"gender"`
	UserID     uuid.UUID `json:"user_id"`
	GroupID    uint      `json:"group_id"`
}
