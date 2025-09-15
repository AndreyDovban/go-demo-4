package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVoult() *Vault {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
}

func (vault *Vault) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
