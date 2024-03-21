package service

import (
	"bufio"
	"fmt"
	"os"
	"pwsd_keeper/model"
	"pwsd_keeper/repository/mysql"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserInfo(userName string) (model.User, error)
}

type Service struct {
	Repo UserRepository
}

func (s Service) createUser() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please insert your name")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("please insert your phone_number")
	scanner.Scan()
	phone_number := scanner.Text()
	fmt.Println("please insert your password")
	scanner.Scan()
	password := scanner.Text()
	fmt.Println("please insert your user_name")
	scanner.Scan()
	user_name := scanner.Text()
	//TODO - check user name is valid and not duplicate and not new_account
	myUser := &model.User{
		UserName:    user_name,
		Name:        name,
		Password:    password,
		PhoneNumber: phone_number,
	}
	if err := s.Repo.CreateUser(myUser); err != nil {
		panic("can not create user " + err.Error())
	} else {
		fmt.Println("your account register successfully\nplease re-run app for login")
		os.Exit(0)
	}

}

func (s Service) loginUser() error {

	return nil
}

func CheckUserStatus() {
	//TODO - read from config file
	appVersion := "0.0.1"
	scanner := bufio.NewScanner(os.Stdin)
	db := mysql.New()
	userService := Service{
		Repo: db,
	}
	fmt.Printf("Welcom to my password keeper app v%s\n", appVersion)
	fmt.Println(`please insert your user_name:
	(if you dont have any account type new_account for create new account)
	`)

	scanner.Scan()
	command := scanner.Text()
	if command == "new_account" {
		userService.createUser()
	} else {
		fmt.Println("login request")
	}

}
