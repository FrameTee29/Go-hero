package domain

type TeacherRepository interface{}

type TeacherUsecase interface {
	GetAllTeacher() (string, error)
}
