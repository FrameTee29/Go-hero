package domain

type AuthUsecase interface {
	CheckSessionTimeout() (SessionResponse, error)
	SignIn() (string, error)
}

type AuthRepository interface {
	CheckSessionTimeout() (bool, error)
}

type SessionResponse struct {
	IsTimeout bool `json:"isTimeout"`
}
