package usecase

import (
	"gohero/bootstrap"
	"gohero/domain"
)

type authUsecase struct {
	app      *bootstrap.Application
	authRepo domain.AuthRepository
}

func NewAuthUsecase(app *bootstrap.Application, authRepo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		app:      app,
		authRepo: authRepo,
	}
}

func (uc *authUsecase) SignIn() (string, error) {
	return "Sign In", nil
}

func (uc *authUsecase) CheckSessionTimeout() (domain.SessionResponse, error) {

	session := domain.SessionResponse{
		IsTimeout: false,
	}

	isTimeout, err := uc.authRepo.CheckSessionTimeout()

	if err != nil {
		return session, err
	}

	session.IsTimeout = isTimeout

	return session, nil
}
