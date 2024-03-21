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
}

func New() *gorm.DB {
	//TODO - read from .env file.
	config := dbconfig{
		user:     "root",
		password: "DB.@nymeet!",
		db_name:  "pswdkeep_db",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.user, config.password, config.db_name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	_, _ = db.DB()
	//TODO - handle error
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	return db
}

/*
func TestDB() *gorm.DB {

}*/

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Password{},
	)
}
