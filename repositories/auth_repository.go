package repositories

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type authRepository struct {
	app *bootstrap.Application
}

func NewAuthRepository(app *bootstrap.Application) domain.AuthRepository {
	return &authRepository{
		app: app,
	}
}

func (*authRepository) CheckSessionTimeout() (bool, error) {
	return false, nil
}
