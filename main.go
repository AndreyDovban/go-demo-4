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
	vault := account.NewVoult(files.NewJsonDb("vault.json"))

	item := promtData([]any{
		"---------------------",
		"1. Создать аккаунт",
		"2. Получить аккаунт",
		"3. Удалить аккаунт",
		"4. Выход",
		"Выберите пункт меню"})
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

func createAccount(vault *account.VaultWithDb) {
	login := promtData([]any{"Введите логин: "})
	password := promtData([]any{"Введите пароль: "})
	url := promtData([]any{"Введите адрес: "})
	var myAccount, err = account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault.AddAccount(*myAccount)

}

func getAccount(vault *account.VaultWithDb) {
	url := promtData([]any{"Введите url для поиска"})
	vault.GetAccountsByUrl(url)
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]any{"Введите url для удаления"})
	vault.DeleteAccountsByUrl(url)
}

func promtData[T []any](prompt T) string {
	var res string

	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

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
