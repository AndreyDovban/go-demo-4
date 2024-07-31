package account

import (
	"demo/password/files"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
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
		color.Red(err.Error())
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
	data, err := vault.ToByte()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json")
}

func (val *Vault) ToByte() ([]byte, error) {
	file, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return file, nil
}
