package mysql

import (
	"errors"
	"gorm.io/gorm"
	"pwsd_keeper/model"
	"pwsd_keeper/service"
)

func (mysql *MYSQLDB) CreateUser(user *model.User) error {
	mysql.database.Create(user)

	return nil
}

func (mysql *MYSQLDB) GetUserInfo(userName string) service.UserLoginResponse {
	var u *model.User
	err := mysql.database.Where(model.User{UserName: userName}).Find(&u).Error
	uLoginRes := service.UserLoginResponse{
		IsUserFind: !errors.Is(err, gorm.ErrRecordNotFound),
		User:       u,
	}

	return uLoginRes
}
