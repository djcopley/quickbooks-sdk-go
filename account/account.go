package account

import (
	"github.com/djcopley/quickbooks-sdk-go"
	"time"
)

type CurrencyRef struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MetaData struct {
	CreateTime      time.Time `json:"CreateTime,omitempty"`
	LastUpdatedTime time.Time `json:"LastUpdatedTime,omitempty"`
}

type Account struct {
	ID                            string      `json:"Id,omitempty"`
	Name                          string      `json:"Name,omitempty"`
	SyncToken                     string      `json:"SyncToken,omitempty"`
	CurrencyRef                   CurrencyRef `json:"CurrencyRef,omitempty"`
	FullyQualifiedName            string      `json:"FullyQualifiedName,omitempty"`
	Domain                        string      `json:"domain,omitempty"`
	Classification                string      `json:"Classification,omitempty"`
	AccountSubType                string      `json:"AccountSubType,omitempty"`
	CurrentBalanceWithSubAccounts float64     `json:"CurrentBalanceWithSubAccounts,omitempty"`
	Sparse                        bool        `json:"sparse,omitempty"`
	MetaData                      MetaData    `json:"MetaData,omitempty"`
	AccountType                   string      `json:"AccountType,omitempty"`
	CurrentBalance                float64     `json:"CurrentBalance,omitempty"`
	Active                        bool        `json:"Active,omitempty"`
	SubAccount                    bool        `json:"SubAccount,omitempty"`
}

var _ quickbooks.Entity = Account{}

func (a Account) GetName() string {
	return "Account"
}
