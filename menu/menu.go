package menu

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func GetMenu() {
	fmt.Println("___Менеджер паролей___")
loop:
	for {
		res := ""
		fmt.Printf(`Выберите пункт меню:
1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход
`)
		fmt.Scan(&res)
		switch res {
		case "1":
			createAccount()
		case "2":
			searchAccount()
		case "3":
			deleteAccount()
		default:
			fmt.Println("Exit")
			break loop
		}

	}

}

func createAccount() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func searchAccount() {
	file, err := files.ReadFile("data.json")
	if err != nil {
		fmt.Println("Не удалось прочитать файл")
		return
	}
	fmt.Println(string(file))
}

func deleteAccount() {
	fmt.Println("Delete")
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
