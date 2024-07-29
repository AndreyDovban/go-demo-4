package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWhithTime struct {
	createAt time.Time
	updateAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	arr := []rune("asdfghjklpiuyrewqzxcvbnm")
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = arr[rand.IntN(len(arr))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*accountWhithTime, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY_STING")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWhithTime{
		createAt: time.Now(), updateAt: time.Now(), account: account{login: login, password: password, url: urlString},
	}

	if password == "" {
		newAcc.account.generatePassword(10)
	}

	return newAcc, nil
}
func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.outputPassword()

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
