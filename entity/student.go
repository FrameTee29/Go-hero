package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Grade     string `gorm:"column:grade"`
}
