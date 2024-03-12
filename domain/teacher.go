package domain

import entity "gohero/entities"

type TeacherRepository interface {
	FindOne(id int) (entity.Teacher, error)
}

type TeacherUsecase interface {
	GetAllTeacher() (string, error)
	GetTeacherById(id int) (entity.Teacher, error)
}
