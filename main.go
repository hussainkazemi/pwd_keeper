package main

import (
	"pwsd_keeper/handler"
)

func main() {
	db := handler.InitDB()
	handler.MainMenu(db)
}
