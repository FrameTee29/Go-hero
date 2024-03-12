package repositories

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type teacherRepository struct {
	app *bootstrap.Application
}

func NewTeacherRepository(app *bootstrap.Application) domain.TeacherRepository {
	return &teacherRepository{
		app: app,
	}
}
