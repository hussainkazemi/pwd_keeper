package main

import (
	"fmt"
	"pwsd_keeper/repository/mysql"
)

func main() {
	fmt.Println("Welcom to first my CLI app")

	db := mysql.New()
	//TODO - run auto migrate out of main
	mysql.AutoMigrate(db)

	fmt.Println("Auto migrate successfully")
}
