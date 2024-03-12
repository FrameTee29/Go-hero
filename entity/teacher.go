package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Level     string `gorm:"column:level"`
}
