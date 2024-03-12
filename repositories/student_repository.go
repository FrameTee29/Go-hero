package repositories

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type studentRepository struct {
	app *bootstrap.Application
}

func NewStudentRepository(app *bootstrap.Application) domain.StudentRepository {
	return &studentRepository{
		app: app,
	}
}
