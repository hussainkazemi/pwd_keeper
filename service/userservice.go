package service

import (
	"pwsd_keeper/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserInfo(userName string) (model.User, error)
}

type Service struct {
	Repo UserRepository
}

func (s Service) CreateUser(user *model.User) error {
	err := s.Repo.CreateUser(user)

	return err
}

/*
func UserLogin(info UserLoginRequest) (UserResponse, error) {

}
*/
