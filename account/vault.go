package account

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVoult(db Db) *VaultWithDb {

	bytes, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{

				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault

	err = json.Unmarshal(bytes, &vault)
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{

				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)

	vault.Save()

	fmt.Println("Аккаунт успешно добавлен")
	acc.Output()
}

func (vault *VaultWithDb) GetAccountsByUrl(url string) {
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

func (vault *VaultWithDb) DeleteAccountsByUrl(url string) {
	result := []Account{}
	daletedAcc := []Account{}

	for _, acc := range vault.Accounts {
		if acc.Url != url {
			result = append(result, acc)
		} else {
			daletedAcc = append(daletedAcc, acc)
		}
	}

	if len(daletedAcc) == 0 {
		fmt.Println("Нет аккаунтов для удаления")
		return
	} else {
		fmt.Println("Удалены следующие аккаунты")
		vault.Accounts = result
		vault.Save()
		for _, acc := range daletedAcc {
			acc.Output()
		}
	}

}

func (vault *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (vault *VaultWithDb) Save() {
	vault.UpdatedAt = time.Now()
	bytes, err := vault.Vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	vault.db.Write(bytes)
}
