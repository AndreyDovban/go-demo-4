package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")

	files.WriteFile("Hello golang", "./data.txt")
	files.ReadFile()

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	myAcc.OutputPassword()

}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
