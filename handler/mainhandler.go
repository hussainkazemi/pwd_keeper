package handler

import (
	"bufio"
	"fmt"
	"os"
	"pwsd_keeper/pkg/utility"
	"pwsd_keeper/repository/mysql"
	"strconv"
)

// TODO: read app version from config file
const (
	APP_VERSION string = "0.0.1"
)

func MainMenu(db *mysql.MYSQLDB) {
	utility.ClearScreen()
	fmt.Printf("Welcom to my password keeper app v%s\n", APP_VERSION)
	fmt.Println(`please select once:
1) Login 
2) Register new account
press enter for Exit`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	n, err := strconv.Atoi(command)
	if err != nil {
		Exit()
	}
	switch n {
	case 1:
		LoginUser(db)
	case 2:
		RegisterUser(db)
	default:
		fmt.Println("your input incorrect")
		Exit()
	}
}

func InitDB() *mysql.MYSQLDB {
	db := *mysql.New()

	return &db
}

func Exit() {
	//utility.ClearScreen()
	fmt.Println("Have nice time ")
	os.Exit(0)
}
