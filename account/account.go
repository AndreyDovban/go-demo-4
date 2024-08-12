package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"
)

type Account struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Url      string    `json:"url"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
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
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	field, _ := reflect.TypeOf(acc).Elem().FieldByName("Login")
	fmt.Println("!!! ", string(field.Tag))

	if password == "" {
		acc.GeneratePassword(12)
	}

	return acc, nil
}
