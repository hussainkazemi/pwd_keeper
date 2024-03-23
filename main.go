package main

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
	"strconv"
	"syscall"
	"time"
)

// TODO: read app version from config file
const (
	APP_VERSION string = "0.0.1"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	userService := initDB()
	mainMenu(scanner, userService)
}

func initDB() service.Service {
	db := mysql.New()
	userService := service.Service{
		Repo: db,
	}
	return userService
}

func mainMenu(scanner *bufio.Scanner, userService service.Service) {
	utility.ClearScreen()
	fmt.Printf("Welcom to my password keeper app v%s\n", APP_VERSION)
	fmt.Println(`please select once:
1) Login 
2) Register new account
press enter for exit`)

	scanner.Scan()
	command := scanner.Text()
	n, err := strconv.Atoi(command)
	if err != nil {
		exit()
	}
	switch n {
	case 1:
		loginUser(scanner, userService)
	case 2:
		registerUser(scanner, userService)
	default:
		fmt.Println("your input incorrect")
		exit()
	}
}

func registerUser(scanner *bufio.Scanner, userService service.Service) {
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

	if err := userService.CreateUser(myUser); err != nil {
		panic("can not create account " + err.Error())
	}
	fmt.Println("your account register successfully\nplease re-run app for login")
	exit()
}

func loginUser(scanner *bufio.Scanner, userService service.Service) {
	isUserLogin := false
	cnt := 0
	for !isUserLogin {
		utility.ClearScreen()
		fmt.Println("please insert your user_name ")
		scanner.Scan()
		userName := scanner.Text()
		fmt.Println("please insert your password (not echo)")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic("can not read password")
		}
		password := string(bytePassword)

		uLoginRes := userService.LoginUser(userName)
		if !uLoginRes.IsUserFind || uLoginRes.User.Password != Ppassowrd.GetMD5Hash(password) {
			cnt += 1
			if cnt == 3 {
				fmt.Println("you try too many ")
				exit()
			}
			fmt.Println("user name not found or password incorrect.\ntry again")
			time.Sleep(time.Second * 2)
		} else {
			isUserLogin = true
		}
	}
	passwordMenu(scanner)

}

func passwordMenu(scanner *bufio.Scanner) {
	utility.ClearScreen()
	fmt.Println(`please select once: 
1) show label list
2) search label
3) update password
4) remove password
5) add new password
6) generate random password
press any key for exit`)
	scanner.Scan()
	command := scanner.Text()
	switch command {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
		
	default:
		fmt.Println("please insert correct number ")
		time.Sleep(time.Second * 2)
		passwordMenu(scanner)
	}
}

func exit() {
	//utility.ClearScreen()
	fmt.Println("Have nice time ")
	os.Exit(0)
}
