package mysql

import (
	"pwsd_keeper/model"
)

func (mysql *MYSQLDB) CreateUser(user *model.User) error {
	mysql.database.Create(user)

	return nil
}

func (mysql *MYSQLDB) GetUserInfo(userName string) (model.User, error) {
	var u *model.User
	mysql.database.Where(model.User{UserName: userName}).First(&u)

	return *u, nil
}
