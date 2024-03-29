package service

import (
	"pwsd_keeper/model"
	"pwsd_keeper/pkg/password"
	"pwsd_keeper/repository/redis"
	"strconv"
)

const (
	KEY = "UserId"
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

func GetCurrentUserId() uint32 {
	ctx, rdb := redis.New()

	userIdStr, err := rdb.Get(ctx, KEY).Result()
	if err != nil {
		panic("can not read current user id")
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		panic("current user id is not digits")
	}
	return uint32(userId)
}

func SetCurrentUserId(userId uint32) {
	ctx, rdb := redis.New()

	err := rdb.Set(ctx, KEY, userId, 0).Err()
	if err != nil {
		panic("can not set current user id ")
	}
}
