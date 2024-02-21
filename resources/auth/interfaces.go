package auth

type ServiceInterface interface {
	GetAllUsers() []User
}
