package main

import (
	"app-demo-4/account"
	"app-demo-4/files"
	"fmt"
)

func main() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите адрес: ")
	var account, err = account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	account.Output()

	files.ReadFile()
	files.WriteFile("./file.json", "hello")

}

func promtData(text string) string {
	var res string
	fmt.Println(text)
	fmt.Scanln(&res)
	return res
}
