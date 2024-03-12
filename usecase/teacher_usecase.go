package usecase

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type teacherUsecase struct {
	app               *bootstrap.Application
	teacherRepository domain.TeacherRepository
}

func NewTeacherUsecase(app *bootstrap.Application, teacherRepository domain.TeacherRepository) domain.TeacherUsecase {
	return &teacherUsecase{
		app:               app,
		teacherRepository: teacherRepository,
	}
}

func (uc *teacherUsecase) GetAllTeacher() (string, error) {
	return "All Teacher Ja", nil
}

func (uc *teacherUsecase) GetTeacherById(id int) (domain.Teacher, error) {
	teacher, err := uc.teacherRepository.FindOne(id)

	if err != nil {
		return teacher, err
	}

	return teacher, nil
}
