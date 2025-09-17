package account

import (
	"app-demo-4/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVoult() *Vault {
	bytes, err := files.ReadFile("vault.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault

	err = json.Unmarshal(bytes, &vault)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	vault.Save()

	fmt.Println("Аккаунт успешно добавлен")
	acc.Output()
}

func (vault *Vault) GetAccountsByUrl(url string) {
	result := []Account{}

	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			result = append(result, acc)
		}
	}

	if len(result) == 0 {
		fmt.Println("Не найденно аккаутов с таким url")
	} else {
		for _, acc := range result {
			acc.Output()
		}
	}
}

func (vault *Vault) DeleteAccountsByUrl(url string) {
	result := []Account{}
	daletedAcc := []Account{}

	for _, acc := range vault.Accounts {
		if acc.Url != url {
			result = append(result, acc)
		} else {
			daletedAcc = append(daletedAcc, acc)
		}
	}

	vault.Accounts = result
	vault.UpdatedAt = time.Now()

	if len(daletedAcc) == 0 {
		fmt.Println("Нет аккаунтов для удаления")
		return
	} else {
		fmt.Println("Удалены следующие аккаунты")
		for _, acc := range daletedAcc {
			acc.Output()
		}
	}

	vault.Save()

}

func (vault *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (vault *Vault) Save() {
	bytes, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	files.WriteFile("vault.json", bytes)
}
