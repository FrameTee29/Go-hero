package domain

import "time"

type TeacherRepository interface {
	FindOne(id int) (Teacher, error)
}

type TeacherUsecase interface {
	GetAllTeacher() (string, error)
	GetTeacherById(id int) (Teacher, error)
}

type Teacher struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
