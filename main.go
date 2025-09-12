package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Account struct {
	Login    string
	Password string
	Url      string
}

func (acc *Account) output() {
	fmt.Println("Login: ", acc.Login)
	fmt.Println("Password: ", acc.Password)
	fmt.Println("Url: ", acc.Url)
	fmt.Println()

}

func (acc *Account) generatePassword(n int) {
	runes := []rune("qwertyuiopasdfghjklzxcvbnm123456789ASDFGHJKLZXCVBNMQWERTYUIOP")
	pass := make([]string, 0, n)

	for range n {
		pass = append(pass, string(runes[rand.IntN(len(runes))]))
	}
	acc.Password = strings.Join(pass, "")

}

func main() {
	var data Account
	var data2 Account

	promtData(&data)
	// promtData(&data2)

	data.output()
	data2.output()

}

func promtData(acc *Account) {

	fmt.Println("----------------------")
	fmt.Println("Введите логин")
	fmt.Scan(&acc.Login)
	fmt.Println("Генерация пароля")
	acc.generatePassword(10)
	fmt.Println("Введите адрес")
	fmt.Scan(&acc.Url)
	fmt.Println("----------------------")

}
