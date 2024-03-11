package domain

type AuthUsecase interface {
	SignIn() (string, error)
}
