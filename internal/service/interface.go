package service

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	ValidateToken(token string) (uint, error)
}
