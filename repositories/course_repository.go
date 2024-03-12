package repositories

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type courseRepository struct {
	app *bootstrap.Application
}

func NewCourseRepository(app *bootstrap.Application) domain.CourseRepository {
	return &courseRepository{
		app: app,
	}
}
