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
	db := files.NewJsonDb("data.json")
	data, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdateAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)
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
	vault.save()

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

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDelete := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDelete = true
	}

	vault.Accounts = accounts
	vault.save()

	return isDelete
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return file, nil
}

func (vault *Vault) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	db := files.NewJsonDb("data.json")
	db.Write(data)
}
