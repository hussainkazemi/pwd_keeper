package mysql

import (
	"pwsd_keeper/model"
)

func (mysql *MYSQLDB) CreatePassword(password *model.Password) error {
	err := mysql.database.Create(password).Error

	return err
}
