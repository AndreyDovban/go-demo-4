package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"time"
)

type Accounts []Account

type Account struct {
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
}

var runes = []rune("qwertyuiopasdfghjklzxcvbnm123456789ASDFGHJKLZXCVBNMQWERTYUIOP")

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("НЕ ПЕРЕДАН ЛОГИН")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Password:  password,
		Login:     login,
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

func (acc *Account) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(acc, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
