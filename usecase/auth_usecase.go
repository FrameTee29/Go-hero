package usecase

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type authUsecase struct {
	app *bootstrap.Application
}

func NewAuthUsecase(app *bootstrap.Application) domain.AuthUsecase {
	return &authUsecase{
		app: app,
	}
}

func (uc *authUsecase) SignIn() (string, error) {
	return "Sign In", nil
}