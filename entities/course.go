package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Max     int    `gorm:"column:max"`
	Min     int    `gorm:"column:min"`
	Subject int    `gorm:"column:subject"`

	TeacherID uint      `gorm:"column:teacher_id"`
	Students  []Student `gorm:"many2many:student_courses;"`
}
