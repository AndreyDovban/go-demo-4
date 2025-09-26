package account

import (
	"app-demo-4/encrypter"

	"encoding/json"
	"fmt"
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
	db  Db
	enc encrypter.Encrypter
}

func NewVoult(db Db, enc encrypter.Encrypter) *VaultWithDb {

	bytes, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{

				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	var vault Vault

	enc.Decrypt(bytes)

	err = json.Unmarshal(bytes, &vault)
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{

				Accounts:  []Account{},
				UpdatedAt: time.Now(),
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
	vault.Accounts = append(vault.Accounts, acc)

	vault.Save()

	fmt.Println("Аккаунт успешно добавлен")
	acc.Output()
}

func (vault *VaultWithDb) GetAccounts(url string, checkUrl func(Account, string) bool) []Account {
	result := []Account{}

	for _, acc := range vault.Accounts {
		if checkUrl(acc, url) {
			result = append(result, acc)
		}
	}

	return result

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
	encrypt_bytes := vault.enc.Encrypt(bytes)
	vault.db.Write(encrypt_bytes)
	// vault.db.Write(bytes)
}
