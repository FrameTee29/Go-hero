package domain

type AuthUsecase interface {
	CheckSession() (bool, error)
	SignIn() (string, error)
}
