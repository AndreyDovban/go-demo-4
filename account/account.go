package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
)

type Account struct {
	Login    string
	Password string
	Url      string
}

type AccountWithTimeStamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Account
}

var runes = []rune("qwertyuiopasdfghjklzxcvbnm123456789ASDFGHJKLZXCVBNMQWERTYUIOP")

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("НЕ ПЕРЕДАН ЛОГИН")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	newAcc := &AccountWithTimeStamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Account: Account{
			Url:      urlString,
			Password: password,
			Login:    login,
		},
	}
	if password == "" {
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("НЕ ПЕРЕДАН ЛОГИН")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	newAcc := &Account{
		Url:      urlString,
		Password: password,
		Login:    login,
	}
	if password == "" {
		newAcc.generatePassword(10)
	}
	return newAcc, nil
}

func (acc *Account) Output() {
	fmt.Println("Login: ", acc.Login)
	fmt.Println("Password: ", acc.Password)
	fmt.Println("Url: ", acc.Url)
	fmt.Println()

}

func (acc *Account) generatePassword(n int) {

	pass := make([]string, 0, n)

	for range n {
		pass = append(pass, string(runes[rand.IntN(len(runes))]))
	}
	acc.Password = strings.Join(pass, "")

}
