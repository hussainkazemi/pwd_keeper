package main

import (
	"fmt"
	"pwsd_keeper/model"
	"pwsd_keeper/repository/mysql"
	"pwsd_keeper/service"
)

func main() {
	//create db:
	/*if err := mysql.CreateDBIfNotExist(); err != nil {
		panic(err)
	}
	fmt.Println("Database created successfully")
	*/
	fmt.Println("Welcom to first my CLI app")

	mysqldb := mysql.New()
	//TODO - run auto migrate out of main
	//mysql.AutoMigrate(mysqldb)
	//fmt.Println("auto migrate successfully")

	myUser := model.User{
		UserName:    "HussainKazemi",
		Name:        "hussain",
		PhoneNumber: "09133321251",
		Password:    "123456",
	}
	loginRequest := service.LoginRequest{
		User:  myUser,
		Store: mysqldb,
	}
	if err := loginRequest.CreateUser(); err != nil {
		panic(err)
	}
	fmt.Println("user add successfully")
}
