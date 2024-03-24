package handler

import (
	"bufio"
	"fmt"
	"os"
	"pwsd_keeper/model"
	"pwsd_keeper/pkg/password"
	"pwsd_keeper/pkg/utility"
	"pwsd_keeper/service"
	"strconv"
	"time"
)

func PasswordMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	utility.ClearScreen()
	fmt.Println(`please select once: 
1) show label list
2) search label
3) update password
4) remove password
5) add new password
6) generate random password
insert exit for exit`)
	scanner.Scan()
	command := scanner.Text()
	switch command {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
		GenerateRandomPassword()
	case "exit":
		Exit()
	default:
		fmt.Println("please insert correct number ")
		time.Sleep(time.Second * 2)
		PasswordMenu()
	}
}

// GenerateRandomPassword create a random strong password for your label.
func GenerateRandomPassword() {
	utility.ClearScreen()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("how many letter (a,b,c,...)? ")
	scanner.Scan()
	letter, _ := strconv.Atoi(scanner.Text())
	fmt.Println("how many number (A,B,C,...)? ")
	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	fmt.Println("how many number (1,2,3,...)? ")
	scanner.Scan()
	capital, _ := strconv.Atoi(scanner.Text())
	fmt.Println("how many signs (!,@,#,$,...)? ")
	scanner.Scan()
	signs, _ := strconv.Atoi(scanner.Text())

	rPasword := model.RandomPassword{
		uint8(letter),
		uint8(capital),
		uint8(number),
		uint8(signs),
		uint8(letter + capital + number + signs),
	}
	isGood := false
	var pwd string
	for !isGood {
		pwd = password.GenerateRandomPassword(rPasword)
		fmt.Printf("is this good password (y/n )? \n %s\n", pwd)
		scanner.Scan()
		ans := scanner.Text()
		if ans == "y" {
			isGood = true
		}
	}

	fmt.Println("do you want to add a label for this password (y/n)? ")

	scanner.Scan()
	ans := scanner.Text()
	if ans == "y" {
		db := InitDB()
		passwordService := service.PasswordService{
			Repo: db,
		}
		fmt.Println("insert label ")
		scanner.Scan()
		label := scanner.Text()
		pModel := model.Password{
			Label:    label,
			Password: pwd,
			UserId:   8,
		}
		err := passwordService.CreatePassword(&pModel)
		if err != nil {
			fmt.Printf("can not add new password in database %s", err.Error())
			time.Sleep(time.Second * 2)
			Exit()
		}
		fmt.Println("your password add successfully")
		time.Sleep(time.Second * 2)
	}
	PasswordMenu()
}
