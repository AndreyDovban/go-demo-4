package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {

	fmt.Println(generatePassword(10))

	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login: login, url: url, password: password,
	}

	outputPassword(&myAccount)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	arr := []rune("asdfghjklpiuyrewqzxcvbnm")
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = arr[rand.IntN(len(arr))]
	}

	return string(res)
}
