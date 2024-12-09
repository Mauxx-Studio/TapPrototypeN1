package models

import (
	"time"

	"gorm.io/gorm"
)

type DaylyAttendance struct {
	gorm.Model

	PersonID          uint
	GroupID           uint
	Day               time.Time
	EntryTime         time.Time
	LeavingTime       time.Time
	WorkSchedule      uint
	Present           bool
	Late              bool
	Retired           bool
	WithNotice        bool
	WithJustification bool
	WorkedTime        float32
	ExtraTime50       float32
	ExtraTime100      float32
}
