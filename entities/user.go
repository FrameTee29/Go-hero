package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"column:first_name"`
}
