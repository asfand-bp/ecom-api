package auth

import "ecom/utils"

type ServiceInterface interface {
	GetAllUsers() ([]User, *utils.Error)
	GetUserByID(id string) (*User, *utils.Error)
	CreateUser(body User) (*User, *utils.Error)
	UpdateUser(id string, body UserUpdateBody) (*User, *utils.Error)
	DeleteUser(id string) (map[string]string, *utils.Error)
}
