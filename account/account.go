package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type Account struct {
	Login    string
	Password string
	Url      string
}

func (acc *Account) OutputPassword() {
	fmt.Println(*acc)
}

func (acc *Account) GeneratePassword(n int) {
	runes := []rune("asdfghjklzxcvbnmqwertyuiop")
	res := make([]rune, n)

	for i := range res {
		res[i] = runes[rand.IntN(len(runes))]
	}

	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY_STRING")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	acc := &Account{
		Login:    login,
		Password: password,
		Url:      urlString,
	}
	if password == "" {
		acc.GeneratePassword(12)
	}

	return acc, nil
}
