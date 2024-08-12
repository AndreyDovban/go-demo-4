package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts []Account `json:"acccounts"`
	UpdateAt time.Time `json:"updateAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err.Error())
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) FindAccountByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, nil
}
