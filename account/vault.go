package account

import (
	"demo/password/encrypter"
	"demo/password/output"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdateAt time.Time `json:"updateAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdateAt: time.Now(),
			},
			db:  db,
			enc: enc,
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
			db:  db,
			enc: enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Vault.Accounts = append(vault.Vault.Accounts, acc)
	vault.save()

}

func (vault *Vault) FindAccounts(url string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, url)
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
		output.Error(err.Error())
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.Error(err.Error())
	}

	vault.db.Write(data)
}
