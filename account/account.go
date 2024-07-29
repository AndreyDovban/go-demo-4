package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type account struct {
	Login    string
	Password string
	Url      string
	CreateAt time.Time
	UpdateAt time.Time
}

func (acc *account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *account) generatePassword(n int) {
	arr := []rune("asdfghjklpiuyrewqzxcvbnm")
	res := make([]rune, n)
	for i := 0; i < n; i++ {
		res[i] = arr[rand.IntN(len(arr))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("LOGIN_EMPTY_STING")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &account{
		CreateAt: time.Now(), UpdateAt: time.Now(), Login: login, Password: password, Url: urlString,
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
