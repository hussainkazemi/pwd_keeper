package service

import (
	"pwsd_keeper/model"
	"pwsd_keeper/pkg/password"
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
	user.Password = password.GetMD5Hash(user.Password)
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
