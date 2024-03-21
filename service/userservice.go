package service

import (
	"pwsd_keeper/model"
)

type Store interface {
	CreateUser(user model.User) error
	GetUserInfo(userName string) (model.User, error)
}

type LoginRequest struct {
	User  model.User
	Store Store
}

func (request LoginRequest) CreateUser() error {
	err := request.Store.CreateUser(request.User)

	return err
}

/*
func UserLogin(info UserLoginRequest) (UserResponse, error) {

}
*/
