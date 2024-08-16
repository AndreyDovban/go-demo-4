package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.Error("Не удалось найти env файл")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	getMenu(vault)
}

func getMenu(vault *account.VaultWithDb) {
	var variant string

loop:
	for {
		variant = promtData([]string{
			"___ Менеджер паролей ___",
			"1. Создать аккаунт",
			"2. Найти аккаунт по Url",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход", "Выберите вариант",
		})

		menuFunc := menu[variant]
		if menuFunc == nil {
			fmt.Println("Exit")
			break loop
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData([]string{"Введите логин: "})
	password := promtData([]string{"Введите пароль: "})
	url := promtData([]string{"Введите URL: "})

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		output.Error(err.Error())
		return
	}

	vault.AddAccount(*myAcc)

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для поиска: "})
	accounts := vault.FindAccounts(url, checkUrl)
	if len(accounts) == 0 {
		fmt.Println("Не найдено аккаунта с таким URL")
		return
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	logun := promtData([]string{"Введите Login для поиска: "})
	accounts := vault.FindAccounts(logun, checkLogin)
	if len(accounts) == 0 {
		fmt.Println("Не найдено аккаунта с таким URL")
		return
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func checkUrl(acc account.Account, url string) bool {
	return strings.Contains(acc.Url, url)
}

func checkLogin(acc account.Account, url string) bool {
	return strings.Contains(acc.Login, url)
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для  удаления: "})
	isDelete := vault.DeleteAccountByUrl(url)
	if isDelete {
		fmt.Println("Удалено")
	} else {
		output.Error("Ненайдено")
	}

}

func promtData[T any](prompt []T) string {
	for i, val := range prompt {
		if i == len(prompt)-1 {
			fmt.Print(val, " : ")
			break
		}
		fmt.Println(val)

	}
	var res string
	fmt.Scanln(&res)
	return res

}
