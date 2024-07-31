package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
}

func (acc *Account) generatePassword(n int) {
	arr := []rune("asdfghjklpiuyrewqzxcvbnm")
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = arr[rand.IntN(len(arr))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY_STING")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreateAt: time.Now(), UpdateAt: time.Now(), Login: login, Password: password, Url: urlString,
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
