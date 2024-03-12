package repositories

import (
	"gohero/bootstrap"
	"gohero/domain"
	entity "gohero/entities"
)

type teacherRepository struct {
	app *bootstrap.Application
}

func NewTeacherRepository(app *bootstrap.Application) domain.TeacherRepository {
	return &teacherRepository{
		app: app,
	}
}

func (r *teacherRepository) FindOne(id int) (entity.Teacher, error) {
	db := r.app.DB
	driverDb := db.DriverDb

	teacher := entity.Teacher{}

	result := driverDb.First(&teacher, id)

	if result.Error != nil {
		return teacher, result.Error
	}

	return teacher, nil
}
