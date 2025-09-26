package main

import (
	"app-demo-4/account"
	"app-demo-4/encrypter"
	"app-demo-4/files"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": getAccountByUrl,
	"3": getAccountByLogin,
	"4": deleteAccount,
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	vault := account.NewVoult(
		files.NewJsonDb("vault"),
		*encrypter.NewEncryptor())

	for {
		item := promtData(
			"---------------------",
			"1. Создать аккаунт",
			"2. Получить аккаунт по Url",
			"3. Получить аккаунт по Login",
			"3. Удалить аккаунт",
			"5. Выход",
			"Выберите пункт меню")
		menuFunc := menu[item]
		if menuFunc == nil {
			break
		}
		menuFunc(vault)
		if !requestToContinue() {
			break
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
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

func getAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите url для поиска")
	accounts := vault.GetAccounts(
		url,
		func(acc account.Account, str string) bool {
			return strings.Contains(acc.Url, str)
		},
	)

	outputResults(&accounts)
}

func getAccountByLogin(vault *account.VaultWithDb) {
	login := promtData("Введите login для поиска")
	accounts := vault.GetAccounts(
		login,
		func(acc account.Account, str string) bool {
			return strings.Contains(acc.Login, str)
		},
	)

	outputResults(&accounts)
}

func outputResults(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		fmt.Println("Не найденно аккаутов")
	} else {
		for _, acc := range *accounts {
			acc.Output()
		}
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData("Введите url для удаления")
	vault.DeleteAccountsByUrl(url)
}

func promtData[T any](prompt ...T) string {
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
