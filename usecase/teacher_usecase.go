package usecase

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type teacherUsecase struct {
	app *bootstrap.Application
}

func NewTeacherUsecase(app *bootstrap.Application) domain.TeacherUsecase {
	return &teacherUsecase{
		app: app,
	}
}

func (uc *teacherUsecase) GetAllTeacher() (string, error) {
	return "All Teacher Ja", nil
}
