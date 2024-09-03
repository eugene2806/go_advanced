package main

import (
	"bookmarks/cmd/cli"
	"bookmarks/cmd/server"
	"fmt"
)

func main() {
	fmt.Println("Приложение для закладок")
	fmt.Println("Выберите врсию")
	fmt.Println("1. CLI")
	fmt.Println("2. REST")
	var opt string
	fmt.Scan(&opt)

	switch opt {
	case "1":
		cli.ConsoleInput()
	case "2":
		server.StartServer()
	}
}
