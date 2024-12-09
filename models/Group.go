package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model

	Name     string
	Manager  uint
	Activity string
}
