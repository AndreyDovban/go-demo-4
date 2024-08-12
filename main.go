package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	vault := account.NewVault()
	getMenu(vault)
}

func getMenu(vault *account.Vault) {
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
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			fmt.Println("Exit")
			break loop

		}
	}

}

func createAccount(vault *account.Vault) {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите URL: ")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	vault.AddAccount(*myAcc)

}

func findAccount(vault *account.Vault) {
	url := promtData("Введите url для поиска: ")
	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("Не найдено аккаунта с таким URL")
		return
	}

	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите url для  удаления: ")
	isDelete := vault.DeleteAccountByUrl(url)
	if isDelete {
		fmt.Println("Удалено")
	} else {
		fmt.Println("Ненайдено")
	}

}

func promtData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
