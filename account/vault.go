package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db *files.JsonDb) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: *db,
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)
	if err != nil {
		fmt.Println(err.Error())
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db: *db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    *db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Vault.Accounts = append(vault.Vault.Accounts, acc)
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

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
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

func (vault *VaultWithDb) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}

	vault.db.Write(data)
}
