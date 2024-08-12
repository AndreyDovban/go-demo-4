package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	getMenu()
}

func getMenu() {
	var variant int

loop:
	for {
		fmt.Println("___ Менеджер паролей ___")
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")

		fmt.Scanln(&variant)

		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			fmt.Println("Exit")
			break loop

		}
	}

}

func createAccount() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAcc)

}

func findAccount() {
	url := promtData("Введите url для поиска: ")

	fmt.Println(url)
}

func deleteAccount() {
	fmt.Println("Delete")
}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
