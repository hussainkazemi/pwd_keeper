package service

import (
	"pwsd_keeper/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserInfo(userName string) UserLoginResponse
}

type UserLoginResponse struct {
	IsUserFind bool
	User       *model.User
}

type Service struct {
	Repo UserRepository
}

type CurrentUser struct {
	user model.User
}

func (s Service) CreateUser(user *model.User) error {
	if err := s.Repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s Service) LoginUser(userName string) UserLoginResponse {
	uLoginReq := s.Repo.GetUserInfo(userName)

	return uLoginReq
}

// TODO - get current user- if user not login return nil
func GetCurrentUser() CurrentUser {

	return CurrentUser{}
}
