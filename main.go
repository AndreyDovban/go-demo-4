package main

import (
	"app-demo-4/account"
	"app-demo-4/files"
	"fmt"
	"strings"
)

func main() {

	for {
		if !getMenu() {
			break
		}
		if !requestToContinue() {
			break
		}
	}
}

func getMenu() bool {
	fmt.Println("---------------------")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Получить аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	item := promtData("Выберите пункт меню: ")
	fmt.Println("---------------------")
	switch {
	case item == "1":
		createAccount()
		return true
	case item == "2":
		getAccount()
		return true
	case item == "3":
		deleteAccount()
		return true
	case item == "4":
		return false
	default:
		return true
	}
}

func createAccount() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите адрес: ")
	var myAccount, err = account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bytes, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	files.WriteFile("./accounts.json", bytes)

	myAccount.Output()

}

func getAccount() {
	fmt.Println("get account")
}

func deleteAccount() {
	fmt.Println("delete account")
}

func promtData(text string) string {
	var res string
	fmt.Println(text)
	fmt.Scanln(&res)
	return res
}

func requestToContinue() bool {
	var res string
	fmt.Println("Продолжить?")
	fmt.Scanln(&res)
	if strings.ToLower(res) == "y" || strings.ToLower(res) == "yes" {
		return true
	}
	return false
}
