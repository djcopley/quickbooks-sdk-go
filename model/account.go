package model

import "time"

type Account struct {
	Account struct {
		FullyQualifiedName string `json:"FullyQualifiedName"`
		Domain             string `json:"domain"`
		Name               string `json:"Name"`
		Classification     string `json:"Classification"`
		AccountSubType     string `json:"AccountSubType"`
		CurrencyRef        struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"CurrencyRef"`
		CurrentBalanceWithSubAccounts int  `json:"CurrentBalanceWithSubAccounts"`
		Sparse                        bool `json:"sparse"`
		MetaData                      struct {
			CreateTime      time.Time `json:"CreateTime"`
			LastUpdatedTime time.Time `json:"LastUpdatedTime"`
		} `json:"MetaData"`
		AccountType    string `json:"AccountType"`
		CurrentBalance int    `json:"CurrentBalance"`
		Active         bool   `json:"Active"`
		SyncToken      string `json:"SyncToken"`
		Id             string `json:"Id"`
		SubAccount     bool   `json:"SubAccount"`
	} `json:"Account"`
	Time time.Time `json:"time"`
}

func (a Account) GetObjectInfo() *QuickbooksObjectInfo {
	return &QuickbooksObjectInfo{
		EntityName: "account",
	}
}
