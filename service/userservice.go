package service

import (
	"pwsd_keeper/model"
)

type UserCreateRequest struct {
	User model.User
}

type UserLoginRequest struct {
	UserName string
	password string
}

type UserResponse struct {
	User model.User
}

/*
func CreateUser(info UserCreateRequest) (UserResponse, error) {

}

func UserLogin(info UserLoginRequest) (UserResponse, error) {

}
*/
