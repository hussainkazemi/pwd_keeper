package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pwsd_keeper/model"
	"pwsd_keeper/pkg/utility"
)

const (
	DB_NAME     = "DB_NAME"
	DB_PORT     = "DB_PORT"
	DB_HOST     = "DB_HOST"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
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

func getDBConfig() dbconfig {
	name, err1 := utility.LoadFromEnv(DB_NAME)
	user, err2 := utility.LoadFromEnv(DB_USER)
	password, err3 := utility.LoadFromEnv(DB_PASSWORD)
	host, err4 := utility.LoadFromEnv(DB_HOST)
	port, err5 := utility.LoadFromEnv(DB_PORT)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		panic("can not connect to db")
	}
	config := dbconfig{
		user:     user,
		password: password,
		db_name:  name,
		host:     host,
		port:     port,
	}
	return config
}

func CreateDBIfNotExist() error {

	config := getDBConfig()

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
	config := getDBConfig()
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
