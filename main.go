package main

import (
	"app-demo-4/account"
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
	vault := account.NewVoult()

	fmt.Println("---------------------")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Получить аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	item := promtData("Выберите пункт меню: ")
	fmt.Println("---------------------")
	switch {
	case item == "1":
		createAccount(vault)
		return true
	case item == "2":
		getAccount(vault)
		return true
	case item == "3":
		deleteAccount(vault)
		return true
	case item == "4":
		return false
	default:
		return true
	}
}

func createAccount(vault *account.Vault) {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите адрес: ")
	var myAccount, err = account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault.AddAccount(*myAccount)

}

func getAccount(vault *account.Vault) {
	url := promtData("Введите url для поиска")
	vault.GetAccountsByUrl(url)
}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите url для удаления")
	vault.DeleteAccountsByUrl(url)
}

func promtData(text string) string {
	var res string
	fmt.Println(text)
	fmt.Scanln(&res)
	return res
}

func requestToContinue() bool {
	var res string
	fmt.Println("Хотите продолжить Y/N?")
	fmt.Scanln(&res)
	if strings.ToLower(res) == "y" || strings.ToLower(res) == "yes" {
		return true
	}
	return false
}
