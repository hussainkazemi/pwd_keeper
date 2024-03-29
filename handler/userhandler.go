package handler

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"pwsd_keeper/model"
	Ppassowrd "pwsd_keeper/pkg/password"
	"pwsd_keeper/pkg/utility"
	"pwsd_keeper/repository/mysql"
	"pwsd_keeper/service"
	"syscall"
	"time"
)

func RegisterUser(db *mysql.MYSQLDB) {
	scanner := bufio.NewScanner(os.Stdin)
	utility.ClearScreen()
	fmt.Println("please insert your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("please insert your phoneNumber: ")
	scanner.Scan()
	phoneNumber := scanner.Text()
	var password string
	for {
		fmt.Println("please insert your password: (not echo)")
		bytePassword1, _ := term.ReadPassword(int(syscall.Stdin))
		fmt.Println("please re-type your password: (not echo)")
		bytePassword2, _ := term.ReadPassword(int(syscall.Stdin))
		if string(bytePassword1) == string(bytePassword2) {
			fmt.Println("password register successfully")
			password = string(bytePassword1)
			break
		} else {
			fmt.Println("passwords are not match please try again")
			time.Sleep(time.Second * 2)
			utility.ClearScreen()
		}
	}
	fmt.Println("please insert your userName: ")
	scanner.Scan()
	userName := scanner.Text()
	//TODO - check user name is valid and not duplicate and not new_account
	myUser := &model.User{
		UserName:    userName,
		Name:        name,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
	userService := service.Service{
		Repo: db,
	}
	if err := userService.CreateUser(myUser); err != nil {
		panic("can not create account " + err.Error())
	}
	fmt.Println("your account register successfully\nplease re-run app for login")
	Exit()
}

func LoginUser(db *mysql.MYSQLDB) {
	isUserLogin := false
	cnt := 0
	for !isUserLogin {
		utility.ClearScreen()
		fmt.Println("please insert your user_name ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userName := scanner.Text()
		fmt.Println("please insert your password (not echo)")
		bytePassword, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			panic("can not read password")
		}
		password := string(bytePassword)
		userService := service.Service{
			Repo: db,
		}
		uLoginRes := userService.LoginUser(userName)
		if !uLoginRes.IsUserFind || uLoginRes.User.Password != Ppassowrd.GetMD5Hash(password) {
			cnt += 1
			if cnt == 3 {
				fmt.Println("you try too many ")
				Exit()
			}
			fmt.Println("user name not found or password incorrect.\ntry again")
			time.Sleep(time.Second * 2)
		} else {
			service.SetCurrentUserId(uLoginRes.User.Id)
			isUserLogin = true
		}
	}
	PasswordMenu()
}
