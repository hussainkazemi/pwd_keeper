package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pwsd_keeper/model"
)

type dbconfig struct {
	user     string
	password string
	db_name  string
	host     string
	port     string
}

type MYSQLDB struct {
	database *gorm.DB
}

func CreateDBIfNotExist() error {
	//TODO - read from .env file.
	config := dbconfig{
		user:     "root",
		password: "DB.@nymeet!",
		db_name:  "pswdkeep_db",
		host:     "localhost",
		port:     "3306",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.user, config.password, config.host, config.port)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.db_name + ";")

	//TODO - handle error every panic
	if err != nil {
		return err
	}

	return nil
}

func New() *MYSQLDB {
	//TODO - read from .env file.
	config := dbconfig{
		user:     "root",
		password: "DB.@nymeet!",
		db_name:  "pswdkeep_db",
		host:     "localhost",
		port:     "3306",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.user, config.password, config.host, config.port, config.db_name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &MYSQLDB{database: db}
}

/*
func TestDB() *gorm.DB {

}*/

func AutoMigrate(mysqlDB *MYSQLDB) {
	err := mysqlDB.database.AutoMigrate(
		&model.User{},
		&model.Password{},
	)
	if err != nil {
		panic(err)
	}
}
