package main

import (
	// "demo/password/account"
	"demo/password/files"
	// "fmt"
)

func main() {
	files.WriteFile("DATA", "file.txt")
	files.ReadFile("file.txt")
	// login := promptData("Введите логин: ")
	// password := promptData("Введите пароль: ")
	// url := promptData("Введите URL: ")

	// myAccount, err := account.NewAccount(login, password, url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// myAccount.OutputPassword()

}

// func promptData(prompt string) string {
// 	fmt.Print(prompt)
// 	var res string
// 	fmt.Scanln(&res)
// 	return res
// }
